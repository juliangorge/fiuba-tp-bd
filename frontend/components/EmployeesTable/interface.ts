interface Employee {
  id: number;
  first_name: string;
  last_name: string;
  position: string;
  department: string;
  hire_date: string;
  salary: number;
}

export const emptyEmployee = {
  first_name: '',
  last_name: '',
  position: '',
  department: '',
  hire_date: '',
  salary: 0,
};

export type { Employee };
