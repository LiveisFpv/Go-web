<template>
  <div class="about">
    <div class="about-container">
      <section class="about-header">
        <h1>О системе</h1>
        <p class="subtitle">Информационная система управления учебным процессом</p>
      </section>

      <section class="about-content">
        <div class="info-card">
          <h2>Назначение системы</h2>
          <p>Система предназначена для автоматизации и оптимизации процессов управления учебным процессом в образовательном учреждении. Она обеспечивает эффективное взаимодействие между всеми участниками образовательного процесса.</p>
        </div>

        <div class="info-card">
          <h2>Основные модули</h2>
          <div class="modules-grid">
            <div class="module-item">
              <h3>Управление студентами</h3>
              <p>Ведение базы данных студентов, управление личными данными и успеваемостью</p>
            </div>
            <div class="module-item">
              <h3>Управление группами</h3>
              <p>Формирование и управление учебными группами, распределение студентов</p>
            </div>
            <div class="module-item">
              <h3>Успеваемость</h3>
              <p>Ведение учета оценок, контроль успеваемости студентов</p>
            </div>
            <div class="module-item">
              <h3>Стипендиальное обеспечение</h3>
              <p>Управление стипендиальными выплатами и бюджетными средствами</p>
            </div>
            <div class="module-item">
              <h3>Достижения</h3>
              <p>Учет достижений студентов и управление категориями достижений</p>
            </div>
            <div class="module-item">
              <h3>Учебные периоды</h3>
              <p>Управление семестрами и учебными периодами</p>
            </div>
          </div>
        </div>

        <div class="info-card">
          <h2>Роли пользователей</h2>
          <div class="roles-grid">
            <div class="role-item">
              <h3>Студент</h3>
              <ul>
                <li>Просмотр личных данных</li>
                <li>Доступ к информации об успеваемости</li>
                <li>Просмотр стипендий</li>
              </ul>
            </div>
            <div class="role-item">
              <h3>Преподаватель</h3>
              <ul>
                <li>Управление оценками</li>
                <li>Просмотр информации о студентах</li>
                <li>Работа с группами</li>
              </ul>
            </div>
            <div class="role-item">
              <h3>Деканат</h3>
              <ul>
                <li>Управление группами</li>
                <li>Контроль успеваемости</li>
                <li>Работа со стипендиями</li>
              </ul>
            </div>
            <div class="role-item">
              <h3>Бухгалтер</h3>
              <ul>
                <li>Управление стипендиями</li>
                <li>Работа с бюджетом</li>
                <li>Финансовая отчетность</li>
              </ul>
            </div>
            <div class="role-item">
              <h3>Администратор</h3>
              <ul>
                <li>Полный доступ к системе</li>
                <li>Управление пользователями</li>
                <li>Настройка системы</li>
              </ul>
            </div>
          </div>
        </div>

        <div class="info-card">
          <h2>Технические возможности</h2>
          <ul class="features-list">
            <li>Современный веб-интерфейс</li>
            <li>Безопасная аутентификация и авторизация</li>
            <li>Защита персональных данных</li>
            <li>Удобная навигация по разделам</li>
            <li>Адаптивный дизайн</li>
            <li>Оптимизированная работа с данными</li>
          </ul>
        </div>
      </section>
    </div>
  </div>
</template>

<style scoped>
.about {
  min-height: 100vh;
  background-color: var(--color-background);
  padding: 2rem 0;
}

.about-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
}

.about-header {
  text-align: center;
  margin-bottom: 3rem;
  padding: 3rem 0;
  background: linear-gradient(135deg, var(--color-background-soft) 0%, var(--color-background) 100%);
  border-radius: 16px;
}

.about-header h1 {
  font-size: 2.5rem;
  color: var(--color-heading);
  margin-bottom: 1rem;
}

.subtitle {
  font-size: 1.2rem;
  color: var(--color-text);
  opacity: 0.8;
}

.about-content {
  display: grid;
  gap: 2rem;
}

.info-card {
  background-color: var(--color-background-soft);
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.info-card h2 {
  color: var(--color-accent);
  font-size: 1.8rem;
  margin-bottom: 1.5rem;
}

.modules-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1.5rem;
}

.module-item {
  background-color: var(--color-background);
  padding: 1.5rem;
  border-radius: 8px;
  border: 1px solid var(--color-border);
}

.module-item h3 {
  color: var(--color-heading);
  font-size: 1.3rem;
  margin-bottom: 0.8rem;
}

.module-item p {
  color: var(--color-text);
  font-size: 0.95rem;
  line-height: 1.5;
}

.roles-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
}

.role-item {
  background-color: var(--color-background);
  padding: 1.5rem;
  border-radius: 8px;
  border: 1px solid var(--color-border);
}

.role-item h3 {
  color: var(--color-heading);
  font-size: 1.3rem;
  margin-bottom: 1rem;
}

.role-item ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.role-item li {
  color: var(--color-text);
  font-size: 0.95rem;
  padding: 0.5rem 0;
  border-bottom: 1px solid var(--color-border);
}

.role-item li:last-child {
  border-bottom: none;
}

.features-list {
  list-style: none;
  padding: 0;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
}

.features-list li {
  padding: 1rem;
  background-color: var(--color-background);
  border-radius: 8px;
  color: var(--color-text);
  border: 1px solid var(--color-border);
}

@media (max-width: 768px) {
  .about-container {
    padding: 0 1rem;
  }

  .about-header {
    padding: 2rem 0;
  }

  .about-header h1 {
    font-size: 2rem;
  }

  .info-card {
    padding: 1.5rem;
  }

  .modules-grid,
  .roles-grid,
  .features-list {
    grid-template-columns: 1fr;
  }
}
</style>
