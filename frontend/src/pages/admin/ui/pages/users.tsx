import AppSidebar from "../components/sidebar";
import Header from "../components/header";
import UserTable from "../components/usertable";
import { User } from '@pages/admin/index';

const users: User[] = [
  { id: 1, name: "Иван Иванов", email: "ivan@example.com", role: "Администратор" },
  { id: 2, name: "Анна Смирнова", email: "anna@example.com", role: "Пользователь" },
];

const UsersPage: React.FC = () => {
  return <UserTable users={users} />;
};

export default UsersPage;
