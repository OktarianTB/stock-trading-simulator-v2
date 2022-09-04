import React from 'react';
import {
  Grid, CardContent, Typography, Card,
} from '@mui/material/';
import clsx from 'clsx';
import Title from '../Template/Title';
import styles from './Search.module.css';

function BodyText({ text }) {
  return (
    <Typography variant="body2" color="inherit" align="center" display="block">
      {text}
    </Typography>
  );
}

function HeaderText({ text }) {
  return (
    <Typography variant="body1" color="inherit" align="center" display="block">
      {text}
    </Typography>
  );
}

function InfoCard({ stockMetadata, price }) {
  return (
    <Grid container spacing={3}>
      <Grid
        item
        xs={12}
        component={Card}
        className={clsx(styles.card, styles.cardBorder)}
      >
        <CardContent>
          <Title>{stockMetadata.name}</Title>
          <Typography variant="body2">{stockMetadata.description}</Typography>
          <Grid container spacing={3} className={styles.addMargin}>
            <Grid item sm={3} xs={4} className={styles.centerGrid}>
              <div className={styles.information}>
                <HeaderText text="Stock Symbol:" />
                <BodyText text={stockMetadata.ticker} />
              </div>
            </Grid>
            <Grid item sm={3} xs={4} className={styles.centerGrid}>
              <div className={styles.information}>
                <HeaderText text="Current Price:" />
                <BodyText text={Math.round(price * 100) / 100} />
              </div>
            </Grid>
            <Grid item sm={3} xs={4} className={styles.centerGrid}>
              <div className={styles.information}>
                <HeaderText text="Exchange:" />
                <BodyText text={stockMetadata.exchangeCode} />
              </div>
            </Grid>
          </Grid>
        </CardContent>
      </Grid>
    </Grid>
  );
}

export default InfoCard;
