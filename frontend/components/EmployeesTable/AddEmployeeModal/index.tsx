'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Employee, emptyEmployee } from '../interface';

interface AddEmployeeModalProps {
  onAdd: (employee: Omit<Employee, 'id'>) => void;
}

export default function AddEmployeeModal({ onAdd }: AddEmployeeModalProps) {
  const [newEmployee, setNewEmployee] =
    useState<Omit<Employee, 'id'>>(emptyEmployee);

  const handleAdd = () => {
    onAdd(newEmployee);
    setNewEmployee(emptyEmployee);
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>Add Employee</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add New Employee</DialogTitle>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="name" className="text-right">
              Name
            </Label>
            <Input
              id="name"
              value={newEmployee.first_name}
              onChange={(e) =>
                setNewEmployee({ ...newEmployee, first_name: e.target.value })
              }
              className="col-span-3"
            />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="position" className="text-right">
              Position
            </Label>
            <Input
              id="position"
              value={newEmployee.position}
              onChange={(e) =>
                setNewEmployee({ ...newEmployee, position: e.target.value })
              }
              className="col-span-3"
            />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="department" className="text-right">
              Department
            </Label>
            <Input
              id="department"
              value={newEmployee.department}
              onChange={(e) =>
                setNewEmployee({
                  ...newEmployee,
                  department: e.target.value,
                })
              }
              className="col-span-3"
            />
          </div>
        </div>
        <Button onClick={handleAdd}>Add Employee</Button>
      </DialogContent>
    </Dialog>
  );
}
