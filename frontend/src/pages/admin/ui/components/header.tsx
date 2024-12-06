import React from "react";
import styles from "../styles/Header.module.scss";

const Header: React.FC = () => (
  <header className={styles.header}>
    <div className={styles.brand}>Админ-панель</div>
  </header>
);

export default Header;
