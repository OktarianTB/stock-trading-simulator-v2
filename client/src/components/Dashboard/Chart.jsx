import React, { useState, useEffect, useContext } from 'react';
import Axios from 'axios';
import UserContext from '../../context/UserContext';
import Title from '../Template/Title.jsx';
import LineChart from '../Template/LineChart';
import config from '../../config/Config';
import styles from './Dashboard.module.css';
import getRandomTicker from '../../utils/Utils';

function Chart() {
  const { userData } = useContext(UserContext);
  const [stockData, setStockData] = useState(undefined);

  const getStockHistoricalData = async (ticker) => {
    const startDate = new Date();
    startDate.setFullYear(startDate.getFullYear() - 3);

    const headers = { Authorization: `Bearer ${userData.accessToken}` };
    const url = `${config.base_url}/api/v1/data?ticker=${ticker}&frequency=monthly&start_at=${startDate.toISOString()}`;

    await Axios.get(
      url,
      { headers },
    ).then((res) => {
      setStockData({
        ticker: res.data.ticker,
        data: res.data.data,
      });
    });
  };

  useEffect(() => {
    const ticker = getRandomTicker();
    getStockHistoricalData(ticker);
  }, []);

  return (
    <div>
      <div style={{ minHeight: '150px' }}>
        {stockData && (
        <div>
          <Title>
            Explore
            {' '}
            {stockData.ticker}
            &apos;s Stock Chart
          </Title>
          <div className={styles.chart}>
            <LineChart
              data={stockData.data}
              ticker={stockData.ticker}
              duration="three years"
            />
          </div>
        </div>
        )}
      </div>
    </div>
  );
}

/**
 * <LineChart
            pastDataPeriod={stockData.data}
            stockInfo={{ ticker: stockData.ticker }}
            duration="3 years"
          />
*/

export default Chart;
