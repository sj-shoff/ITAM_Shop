import React from "react";
import { useRouter } from "next/router";
import styles from "../styles/Sidebar.module.scss";

const AppSidebar: React.FC = () => {
  const router = useRouter();

  const menuItems = [
    { name: "Главная", path: "/" },
    { name: "Товары", path: "/products" },
    { name: "Пользователи", path: "/users" },
    { name: "Статистика", path: "/stats" },
  ];

  return (
    <div className={styles.sidebar}>
      <h3>Меню</h3>
      <nav>
        {menuItems.map((item) => (
          <button
            key={item.path}
            className={styles["menu-item"]}
            onClick={() => router.push(item.path)}
          >
            {item.name}
          </button>
        ))}
      </nav>
    </div>
  );
};

export default AppSidebar;
