'use client';

import { useEffect, useState } from 'react';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { Employee } from './interface';
import AddEmployeeModal from './AddEmployeeModal';
import EditEmployeeModal from './EditEmployeeModal';
import ViewEmployeeModal from './ViewEmployeeModal';
import DeleteEmployeeModal from './DeleteEmployeeModal';

export default function EmployeesTable() {
  const [employees, setEmployees] = useState<Employee[]>([]);
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [editingEmployee, setEditingEmployee] = useState<Employee | null>(null);

  useEffect(() => {
    const fetchMockData = async () => {
      const response = await fetch('/mocks/employees.json');
      const data: Employee[] = await response.json();
      setEmployees(data);
    };

    fetchMockData();
  }, []);

  const addEmployee = (newEmployee: Omit<Employee, 'id'>) => {
    setEmployees([...employees, { ...newEmployee, id: Date.now() }]);
  };

  const updateEmployee = (updatedEmployee: Employee) => {
    setEmployees(
      employees.map((emp) =>
        emp.id === updatedEmployee.id ? updatedEmployee : emp,
      ),
    );
    setEditingEmployee(null);
  };

  const deleteEmployee = (id: number) => {
    setEmployees(employees.filter((emp) => emp.id !== id));
  };

  return (
    <div className="container mx-auto p-4">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">Employee Management</h1>
        <AddEmployeeModal onAdd={(employee) => addEmployee(employee)} />
      </div>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Position</TableHead>
            <TableHead>Department</TableHead>
            <TableHead className="text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {employees.map((employee) => (
            <TableRow key={employee.id}>
              <TableCell>
                {employee.first_name + ' ' + employee.last_name}
              </TableCell>
              <TableCell>{employee.position}</TableCell>
              <TableCell>{employee.department}</TableCell>
              <TableCell className="text-right">
                <ViewEmployeeModal employee={employee} />
                <EditEmployeeModal
                  employee={employee}
                  onEdit={(updatedEmployee) => updateEmployee(updatedEmployee)}
                />
                <DeleteEmployeeModal
                  employee={employee}
                  onDelete={(deletedEmployee) =>
                    deleteEmployee(deletedEmployee.id)
                  }
                />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
