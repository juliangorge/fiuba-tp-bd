import EmployeesTable from '@/components/EmployeesTable';
import Image from 'next/image';

export default function Home() {
  return (
    <div className="items-center min-h-screen sm:pt-0 p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <main className="flex flex-col gap-4 items-center sm:items-start">
        <div className="w-full flex justify-center">
          <Image
            className="dark:invert"
            src="/logo-fiuba.png"
            alt="FIUBA"
            width={330}
            height={128}
            priority
          />
        </div>

        <EmployeesTable />
      </main>
    </div>
  );
}
