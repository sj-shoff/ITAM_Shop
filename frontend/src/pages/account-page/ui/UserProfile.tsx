import React from 'react';
import { Button, Input, Avatar } from '@nextui-org/react';
import styles from './UserProfile.module.scss';

interface UserInfo {
  email: string;
  password: string;
}

const UserProfile: React.FC = () => {
  const userInfo: UserInfo = {
    email: 'user@example.com',
    password: '1234567890000000000',
  };

  const maskPassword = (password: string): string => {
    const visibleLength = Math.ceil(password.length * 0.1);
    return (
      password.slice(0, visibleLength) +
      '*'.repeat(password.length - visibleLength)
    );
  };

  return (
    <div className={styles.container}>
      <div className={styles.card}>
        <Avatar
          className={styles.avatar}
          size="lg"
          style={{ background: 'linear-gradient(45deg, #ff6b6b, #f3a683)' }}
          icon={<span>👤</span>}
        />
        <h2 className={styles.title}>Имя Пользователя</h2>
        <div className={styles.inputGroup}>
          <Input
            readOnly
            label="Почта"
            value={userInfo.email}
            variant="bordered"
            fullWidth
            className={styles.input}
          />
          <Input
            readOnly
            label="Пароль"
            value={maskPassword(userInfo.password)}
            variant="bordered"
            fullWidth
            className={styles.input}
          />
        </div>
        <div className={styles.buttons}>
          <Button size="md" className={styles.button}>
            Корзина
          </Button>
          <Button size="md" className={styles.button}>
            Избранное
          </Button>
        </div>
      </div>
    </div>
  );
};

export default UserProfile;
