import React from 'react';
import { Line } from 'react-chartjs-2';
// eslint-disable-next-line no-unused-vars
import { Chart, Title } from 'chart.js/auto';

const LineChart = ({ data, ticker, duration }) => {
  const formatDate = (date) => {
    const d = new Date(date);
    let month = `${d.getMonth() + 1}`;
    let day = `${d.getDate()}`;

    if (month.length < 2) month = `0${month}`;
    if (day.length < 2) day = `0${day}`;

    return [month, day].join('-');
  };

  const lineChart = data.length > 0 ? (
    <Line
      data={{
        labels: data.map(({ date }) => formatDate(date)),
        datasets: [
          {
            data: data.map(({ price }) => price),
            label: 'Price',
            borderColor: 'rgba(0, 0, 255, 0.5)',
            fill: true,
            backgroundColor: 'rgba(116, 185, 255, 0.2)',
          },
        ],
      }}
      options={{
        maintainAspectRatio: false,
        elements: {
          point: {
            radius: 2,
          },
        },
        plugins: {
          legend: {
            display: false,
          },
          title: {
            display: true,
            text: `Adjusted closing stock price of ${ticker} over the past ${duration}`,
            position: 'bottom',
          },
        },
        layout: {
          padding: {
            left: 20,
            right: 20,
            top: 20,
            bottom: 0,
          },
        },
        animation: {
          duration: 2000,
        },
      }}
    />
  ) : null;

  return lineChart;
};

export default LineChart;
