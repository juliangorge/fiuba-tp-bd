'use client';

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
import { Employee } from '../interface';

interface EditEmployeeModalProps {
  employee: Employee | null;
  onEdit: (employee: Employee) => void;
}

export default function EditEmployeeModal({
  employee,
  onEdit,
}: EditEmployeeModalProps) {
  const handleEdit = (field: keyof Employee, value: string) => {
    if (employee) {
      onEdit({ ...employee, [field]: value });
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="ghost">Edit</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Edit Employee</DialogTitle>
        </DialogHeader>
        {employee && (
          <div className="grid gap-4 py-4">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="edit-name" className="text-right">
                Name
              </Label>
              <Input
                id="edit-name"
                value={employee.first_name}
                onChange={(e) => handleEdit('first_name', e.target.value)}
                className="col-span-3"
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="edit-position" className="text-right">
                Position
              </Label>
              <Input
                id="edit-position"
                value={employee.position}
                onChange={(e) => handleEdit('position', e.target.value)}
                className="col-span-3"
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="edit-department" className="text-right">
                Department
              </Label>
              <Input
                id="edit-department"
                value={employee.department}
                onChange={(e) => handleEdit('department', e.target.value)}
                className="col-span-3"
              />
            </div>
          </div>
        )}
        <Button onClick={() => onEdit(employee!)}>Update Employee</Button>
      </DialogContent>
    </Dialog>
  );
}
