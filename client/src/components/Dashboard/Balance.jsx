import React, { useContext } from 'react';
import { Typography } from '@mui/material/';
import UserContext from '../../context/UserContext';
import Title from '../Template/Title.jsx';
import styles from './Dashboard.module.css';

function Balance({ stockBalance }) {
  const { userData } = useContext(UserContext);

  return (
    <>
      <Title>Current Balance</Title>
      <br />
      <div className={styles.depositContext}>
        <Typography color="textSecondary" align="center">
          Cash Balance:
        </Typography>
        <Typography component="p" variant="h6" align="center">
          $
          {userData ? userData.user.balance.toLocaleString() : '---'}
        </Typography>

        <Typography color="textSecondary" align="center">
          Portfolio Balance:
        </Typography>
        <Typography component="p" variant="h6" align="center" gutterBottom>
          $
          {stockBalance ? stockBalance.toLocaleString() : '---'}
        </Typography>

        <div className={styles.addMargin}>
          <Typography color="textSecondary" align="center">
            Total:
          </Typography>

          <Typography
            component="p"
            variant="h4"
            align="center"
            className={
              Number(userData.user.balance + stockBalance) >= 100000
                ? styles.positive
                : styles.negative
            }
          >
            $
            {userData
              ? (userData.user.balance + stockBalance).toLocaleString()
              : '---'}
          </Typography>
        </div>
      </div>
      <div>
        <Typography color="textSecondary" align="center">
          {new Date().toDateString()}
        </Typography>
      </div>
    </>
  );
}

export default Balance;
