export interface Pagination{
  total: number
  page: number
  page_size: number
}

export interface Sort{
  direction: string
  by: string
}

export interface Filter{
  field: string
  value: string
}
