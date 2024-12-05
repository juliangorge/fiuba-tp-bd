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
      fetch(`http://localhost:8080/employee_tags/${employee.id}`)
        .then((res) => res.json())
        .then((data) => {
          if (data.tags) {
            setTags(data.tags);
          } else {
            setTags([]);
          }
        })
        .catch((err) => console.error('Failed to fetch tags:', err));
    }
  }, [employee]);

  const addTag = (newTag: string) => {
    if (!newTag.trim()) return;
    fetch(`http://localhost:8080/employee_tags/${employee?.id}/tags/${newTag}`, {
      method: 'POST',
    })
      .then((res) => {
        if (res.ok) {
          setTags((prev) => [...prev, newTag]);
        } else {
          console.error('Failed to add tag');
        }
      })
      .catch(console.error);
  };

  const removeTag = (tagToRemove: string) => {
    fetch(`http://localhost:8080/employee_tags/${employee?.id}/tags/${tagToRemove}`, {
      method: 'DELETE',
    })
      .then((res) => {
        if (res.ok) {
          setTags((prev) => prev.filter((tag) => tag !== tagToRemove));
        } else {
          console.error('Failed to remove tag');
        }
      })
      .catch(console.error);
  };

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
                    <li key={index} className="tag-item flex items-center gap-2">
                      {tag}
                      <button
                        className="text-red-500 hover:text-red-700"
                        onClick={() => removeTag(tag)}
                      >
                        x
                      </button>
                    </li>
                  ))}
                </ul>
              ) : (
                <p>No tags available</p>
              )}
            </div>
            <div className="mt-4">
              <form
                onSubmit={(e) => {
                  e.preventDefault();
                  const form = e.target as HTMLFormElement;
                  const input = form.elements.namedItem('new-tag') as HTMLInputElement;
                  addTag(input.value);
                  input.value = '';
                }}
              >
                <input
                  type="text"
                  name="new-tag"
                  className="border p-2 rounded mr-2"
                  placeholder="Enter tag name"
                />
                <button type="submit" className="text-white bg-blue-500 px-4 py-2 rounded hover:bg-blue-700 text-xl">
                  +
                </button>
              </form>
            </div>
          </div>
        )}
      </DialogContent>
    </Dialog>
  );
}