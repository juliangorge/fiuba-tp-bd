'use client';

import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Label } from '@/components/ui/label';
import { Employee } from '../interface';

interface ViewEmployeeModalProps {
  employee: Employee | null;
}

export default function ViewEmployeeModal({
  employee,
}: ViewEmployeeModalProps) {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="ghost">View</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Employee</DialogTitle>
        </DialogHeader>
        {employee && (
          <div className="grid gap-4 py-4">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="edit-name" className="text-right">
                Name
              </Label>
              {employee.first_name + ' ' + employee.last_name}
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="edit-position" className="text-right">
                Position
              </Label>
              {employee.position}
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="edit-department" className="text-right">
                Department
              </Label>
              {employee.department}
            </div>

            <p>And here we should show no-sql data</p>
          </div>
        )}
      </DialogContent>
    </Dialog>
  );
}
