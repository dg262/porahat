import React from 'react';

// TableColumn definition, which is generic
interface TableColumn<T> {
  title: string;
  render: (item: T) => React.ReactNode;
}

// TableProps definition, using the generic type T
interface TableProps<T> {
  data: T[];
  columns: TableColumn<T>[];
}

// Generic Table component
function Table<T>({ data, columns }: TableProps<T>) {
  return (
    <table>
      <thead>
        <tr>
          {columns.map((col, index) => (
            <th key={index}>{col.title}</th>
          ))}
        </tr>
      </thead>
      <tbody>
        {data.map((item, rowIndex) => (
          <tr key={rowIndex}>
            {columns.map((col, colIndex) => (
              <td key={colIndex}>{col.render(item)}</td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
}

export default Table;
