import React, { useState, useEffect, useContext } from 'react';
import {
  Grid, Box, Card,
} from '@mui/material/';
import Axios from 'axios';
import UserContext from '../../context/UserContext';
import LineChart from '../Template/LineChart';
import Copyright from '../Template/Copyright';
import styles from './Search.module.css';
import InfoCard from './InfoCard';
import PurchaseCard from './PurchaseCard';
import PurchaseModal from './PurchaseModal';
import config from '../../config/Config';

function LineChartCardPastMonth({ data, ticker, duration }) {
  return (
    <Grid
      item
      xs={12}
      sm={7}
      component={Card}
      className={styles.card}
      style={{ minHeight: '350px' }}
    >
      <LineChart
        data={data}
        ticker={ticker}
        duration={duration}
      />
    </Grid>
  );
}

function LineChartCardPastTwoYears({ data, ticker, duration }) {
  return (
    <Grid
      item
      xs={12}
      sm={12}
      component={Card}
      className={styles.card}
      style={{ minHeight: '450px' }}
    >
      <LineChart
        data={data}
        ticker={ticker}
        duration={duration}
      />
    </Grid>
  );
}

function StockCard({ currentStock }) {
  const { userData } = useContext(UserContext);
  const [selected, setSelected] = useState(false);
  const [stockMetadata, setStockMetadata] = useState(undefined);
  const [stockLatestPrice, setStockLatestPrice] = useState(undefined);
  const [stockDataPastMonth, setStockDataPastMonth] = useState(undefined);
  const [stockDataPastTwoYears, setStockDataPastTwoYears] = useState(undefined);

  const getStockMetadata = async () => {
    const headers = { Authorization: `Bearer ${userData.accessToken}` };
    const url = `${config.base_url}/api/v1/metadata?ticker=${currentStock.ticker}`;

    await Axios.get(
      url,
      { headers },
    ).then((res) => {
      setStockMetadata(res.data);
    });
  };

  const getStockDataPastMonth = async () => {
    const startDate = new Date();
    startDate.setMonth(startDate.getMonth() - 1);

    const headers = { Authorization: `Bearer ${userData.accessToken}` };
    const url = `${config.base_url}/api/v1/data?ticker=${currentStock.ticker}&frequency=daily&start_at=${startDate.toISOString()}`;

    await Axios.get(
      url,
      { headers },
    ).then((res) => {
      setStockLatestPrice(res.data.data[0]);
      setStockDataPastMonth(res.data.data);
    });
  };

  const getStockDataPastTwoYears = async () => {
    const startDate = new Date();
    startDate.setFullYear(startDate.getFullYear() - 2);

    const headers = { Authorization: `Bearer ${userData.accessToken}` };
    const url = `${config.base_url}/api/v1/data?ticker=${currentStock.ticker}&frequency=monthly&start_at=${startDate.toISOString()}`;

    await Axios.get(
      url,
      { headers },
    ).then((res) => {
      setStockDataPastTwoYears(res.data.data);
    });
  };

  useEffect(() => {
    getStockMetadata();
    getStockDataPastMonth();
    getStockDataPastTwoYears();
  }, []);

  return (
    <div className={styles.root}>
      {stockMetadata && stockLatestPrice && (
        <InfoCard stockMetadata={stockMetadata} price={stockLatestPrice.price} />
      )}
      {stockDataPastMonth && stockDataPastTwoYears && (
        <div>
          <Grid container spacing={3}>
            <LineChartCardPastTwoYears
              data={stockDataPastTwoYears}
              ticker={currentStock.ticker}
              duration="two years"
            />
          </Grid>
          <Grid container spacing={3}>
            <PurchaseCard
              setSelected={setSelected}
              balance={userData.user.balance}
            />
            <LineChartCardPastMonth
              data={stockDataPastMonth}
              ticker={currentStock.ticker}
              duration="month"
            />
          </Grid>
          <Box pt={4}>
            <Copyright />
          </Box>
          {selected && (
            <PurchaseModal
              stockMetadata={stockMetadata}
              price={stockLatestPrice.price}
              setSelected={setSelected}
            />
          )}
        </div>
      )}
    </div>
  );
}
export default StockCard;
