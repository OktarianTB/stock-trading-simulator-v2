import React, { useState, useEffect, useContext } from 'react';
import { styled } from '@mui/system';
import {
  Box, Container, Grid, Paper, CircularProgress,
} from '@mui/material';
import Axios from 'axios';
import UserContext from '../../context/UserContext';
import styles from '../Template/PageTemplate.module.css';
import Chart from './Chart';
import Balance from './Balance';
import Stocks from './Stocks';
import Transactions from './Transactions';
import Copyright from '../Template/Copyright';
import config from '../../config/Config';

const PaperCard = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(2),
  display: 'flex',
  overflow: 'auto',
  flexDirection: 'column',
}));

function Dashboard() {
  const { userData } = useContext(UserContext);
  const [stocks, setStocks] = useState(null);
  const [transactions, setTransactions] = useState(null);

  const getStocksForUser = async () => {
    const headers = { Authorization: `Bearer ${userData.accessToken}` };

    await Axios.get(
      `${config.base_url}/api/v1/stocks`,
      { headers },
    ).then((res) => {
      setStocks({
        balance: res.data.portfolio_balance,
        stocks: res.data.stocks,
      });
    });
  };

  const getTransactionsForUser = async () => {
    const headers = { Authorization: `Bearer ${userData.accessToken}` };

    await Axios.get(
      `${config.base_url}/api/v1/transactions?page_id=1&page_size=15`,
      { headers },
    ).then((res) => {
      setTransactions(res.data);
    });
  };

  useEffect(() => {
    getStocksForUser();
    getTransactionsForUser();
  }, []);

  if (!stocks || !transactions) {
    return (
      <Container maxWidth="lg" className={styles.container}>
        <Box sx={{ display: 'flex' }} className={styles.spinner}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" className={styles.container}>
      <Grid container spacing={3}>
        {/* Chart */}
        <Grid item xs={12} md={8} lg={9}>
          <PaperCard sx={{ height: '400px' }}>
            <Chart />
          </PaperCard>
        </Grid>
        {/* Balance */}
        <Grid item xs={12} md={4} lg={3}>
          <PaperCard sx={{ height: '400px' }}>
            <Balance stockBalance={stocks.balance} />
          </PaperCard>
        </Grid>
        {/* Stocks */}
        <Grid item xs={12}>
          <PaperCard>
            <Stocks stocks={stocks.stocks} />
          </PaperCard>
        </Grid>
        <Grid item xs={12}>
          <PaperCard>
            <Transactions transactions={transactions} />
          </PaperCard>
        </Grid>
      </Grid>
      <Box pt={4}>
        <Copyright />
      </Box>
    </Container>
  );
}

export default Dashboard;
