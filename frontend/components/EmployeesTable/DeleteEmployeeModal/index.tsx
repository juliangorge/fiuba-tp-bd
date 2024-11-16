'use client';

import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Employee } from '../interface';

interface DeleteEmployeeModalProps {
  employee: Employee | null;
  onDelete: (employee: Employee) => void;
}

export default function DeleteEmployeeModal({
  employee,
  onDelete,
}: DeleteEmployeeModalProps) {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="ghost">Delete</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Delete Employee</DialogTitle>
        </DialogHeader>
        {employee && (
          <>
            <div className="grid gap-4 py-4">
              <div className="grid grid-cols-4 items-center gap-4">
                <p>Are you sure to delete this employee?</p>
              </div>
            </div>
            <Button onClick={() => onDelete(employee!)}>Remove Employee</Button>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
}
