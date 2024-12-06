import React from "react";
import { User } from '@pages/admin/index';

const users: User[] = [
  { id: 1, name: "Иван Иванов", email: "ivan@example.com", role: "Администратор" },
  { id: 2, name: "Анна Смирнова", email: "anna@example.com", role: "Пользователь" },
];

interface UserTableProps {
    users: User[];
  }
  
  const UserTable: React.FC<UserTableProps> = ({ users }) => {
    return (
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Имя</th>
            <th>Email</th>
            <th>Роль</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.id}>
              <td>{user.id}</td>
              <td>{user.name}</td>
              <td>{user.email}</td>
              <td>{user.role}</td>
            </tr>
          ))}
        </tbody>
      </table>
    );
  };
  
  export default UserTable;

