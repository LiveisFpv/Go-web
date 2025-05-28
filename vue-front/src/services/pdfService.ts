import * as pdfMake from 'pdfmake/build/pdfmake';
import * as pdfFonts from 'pdfmake/build/vfs_fonts';
import type { TDocumentDefinitions } from 'pdfmake/interfaces';
import { useAuthStore } from '@/stores/auth';

// Устанавливаем шрифты
(pdfMake as any).vfs = pdfFonts.vfs;

interface TableConfig {
  title: string;
  headers: string[];
  widths: number[];
  getRowData: (item: any) => any[];
}

const tableConfigs: Record<string, TableConfig> = {
  students: {
    title: 'Отчет по студентам',
    headers: ['ID студента', 'ФИО', 'Группа', 'Email'],
    widths: [50, 150, 100, 150],
    getRowData: (student) => [
      student.id_num_student,
      `${student.second_name_student} ${student.first_name_student} ${student.surname_student}`,
      student.name_group,
      student.email_student
    ]
  },
  marks: {
    title: 'Отчет по оценкам',
    headers: ['ID студента', 'ФИО', 'Группа', 'Семестр', 'Предмет', 'Оценка', 'Тип оценки', 'Тип экзамена'],
    widths: [50, 150, 100, 100, 150, 80, 100, 100],
    getRowData: (mark) => [
      mark.id_num_student,
      `${mark.second_name_student} ${mark.first_name_student} ${mark.surname_student}`,
      mark.name_group,
      mark.name_semester,
      mark.lesson_name_mark,
      mark.score_mark,
      mark.type_mark,
      mark.type_exam
    ]
  },
  semesters: {
    title: 'Отчет по семестрам',
    headers: ['Название', 'Дата начала', 'Дата окончания'],
    widths: [100, 100, 100],
    getRowData: (semester) => [
      semester.name_semester,
      semester.date_start_semester,
      semester.date_end_semester
    ]
  },
  scholarships: {
    title: 'Отчет по стипендиям',
    headers: ['Студент', 'ФИО', 'Группа', 'Семестр', 'Размер', 'Тип стипендии'],
    widths: [50, 150, 100, 100, 80, 120],
    getRowData: (scholarship) => [
      scholarship.id_num_student,
      `${scholarship.second_name_student} ${scholarship.first_name_student} ${scholarship.surname_student}`,
      scholarship.name_group,
      scholarship.name_semester,
      scholarship.size_scholarshp,
      scholarship.type_scholarship_budget
    ]
  },
  groups: {
    title: 'Отчет по группам',
    headers: ['Название группы', 'Направление обучения', 'Профиль обучения', 'Дата начала', 'Срок обучения'],
    widths: [100, 150, 150, 100, 100],
    getRowData: (group) => [
      group.name_group,
      group.studies_direction_group,
      group.studies_profile_group,
      group.start_date_group,
      group.studies_period_group
    ]
  },
  categories: {
    title: 'Отчет по категориям',
    headers: ['Тип достижения', 'Баллы'],
    widths: [200, 100],
    getRowData: (category) => [
      category.achivments_type_category,
      category.score_category
    ]
  },
  budgets: {
    title: 'Отчет по бюджетам',
    headers: ['Размер', 'Тип стипендии', 'Семестр'],
    widths: [100, 150, 100],
    getRowData: (budget) => [
      budget.size_budget,
      budget.type_scholarship_budget,
      budget.name_semester
    ]
  },
  achievements: {
    title: 'Отчет по достижениям',
    headers: ['Студент', 'ФИО', 'Группа', 'Название достижения', 'Дата', 'Тип достижения'],
    widths: [50, 150, 100, 150, 100, 150],
    getRowData: (achievement) => [
      achievement.id_num_student,
      `${achievement.second_name_student} ${achievement.first_name_student} ${achievement.surname_student}`,
      achievement.name_group,
      achievement.name_achivement,
      achievement.date_achivement,
      achievement.achivments_type_category
    ]
  }
};

export const pdfService = {
  async generateReport(
    data: any[],
    type: string,
    name: string = 'Отчет',
    creator: string = useAuthStore().email || 'Пользователь',
    organization: string = 'Ниженский Государственный институт',
    note: string = 'Отчет сгенерирован автоматически'
  ) {
    try {
      const config = tableConfigs[type];
      if (!config) {
        throw new Error(`Unknown table type: ${type}`);
      }

      const now = new Date();

      // Определяем структуру документа
      const docDefinition: TDocumentDefinitions = {
        pageOrientation: 'landscape',
        content: [
          // Заголовок
          {
            text: name || config.title,
            style: 'header',
            alignment: 'center',
            margin: [0, 0, 0, 10]
          },
          // Метаданные
          {
            text: [
              { text: 'Дата и время создания: ', bold: true },
              now.toLocaleString(),
              '\n',
              { text: 'Создатель: ', bold: true },
              creator,
              '\n',
              { text: 'Организация: ', bold: true },
              organization,
              '\n',
              { text: 'Примечание: ', bold: true },
              note
            ],
            margin: [0, 0, 0, 10]
          },
          // Таблица
          {
            table: {
              headerRows: 1,
              widths: config.widths,
              body: [
                // Заголовки таблицы
                config.headers,
                // Данные таблицы
                ...data.map(item => config.getRowData(item))
              ]
            },
            layout: {
              fillColor: function(rowIndex: number) {
                return rowIndex === 0 ? '#428bca' : (rowIndex % 2 === 0 ? '#f5f5f5' : null);
              },
              hLineWidth: function() { return 0.5; },
              vLineWidth: function() { return 0.5; },
              hLineColor: function() { return '#aaa'; },
              vLineColor: function() { return '#aaa'; },
              paddingLeft: function() { return 2; },
              paddingRight: function() { return 2; },
              paddingTop: function() { return 1; },
              paddingBottom: function() { return 1; },
            }
          }
        ],
        styles: {
          header: {
            fontSize: 16,
            bold: true,
            margin: [0, 0, 0, 5]
          }
        },
        defaultStyle: {
          fontSize: 8,
          lineHeight: 1.2
        },
        pageSize: 'A4',
        pageMargins: [20, 40, 20, 40],
        footer: function(currentPage: number, pageCount: number) {
          return {
            text: `Страница ${currentPage} из ${pageCount}`,
            alignment: 'center',
            margin: [0, 0, 0, 10]
          };
        }
      };

      // Создаем PDF
      pdfMake.createPdf(docDefinition).download(`${type}_report.pdf`);

      return true;
    } catch (error) {
      console.error('Error generating PDF:', error);
      throw error;
    }
  }
};
