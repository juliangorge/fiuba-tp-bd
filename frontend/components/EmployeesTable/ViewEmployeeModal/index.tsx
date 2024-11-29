'use client';

import { useEffect, useState } from 'react';
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
  const [tags, setTags] = useState<string[]>([]);

  useEffect(() => {
    if (employee) {
      fetch(`/api/employee_tags/${employee.id}`)
        .then((res) => res.json())
        .then((data) => {
          if (data.tags) {
            setTags(data.tags);
          } else {
            setTags([]);
          }
        })
    }
  }, [employee]);

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
            <div>
              <h3 className="font-bold">Tags</h3>
              {tags.length > 0 ? (
                <ul>
                  {tags.map((tag, index) => (
                    <li key={index} className="tag-item">
                      {tag}
                    </li>
                  ))}
                </ul>
              ) : (
                <p>No tags available</p>
              )}
            </div>
          </div>
        )}
      </DialogContent>
    </Dialog>
  );
}
