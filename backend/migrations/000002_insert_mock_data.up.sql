-- Check if the table is empty before inserting mock data
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM employees) THEN
        INSERT INTO employees (first_name, last_name, position, department, hire_date, salary) VALUES
        ('John', 'Doe', 'Software Engineer', 'Engineering', '2022-01-15', 70000),
        ('Jane', 'Smith', 'Product Manager', 'Product', '2021-03-22', 80000),
        ('Alice', 'Johnson', 'Designer', 'Design', '2020-07-30', 65000),
        ('Bob', 'Brown', 'Data Scientist', 'Data', '2019-11-05', 90000);
    END IF;
END $$;