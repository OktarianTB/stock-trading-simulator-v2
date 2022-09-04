import React, { useState } from 'react';
import {
  Link, Table, TableBody, TableCell, TableHead, TableRow,
} from '@mui/material';
import Title from '../Template/Title.jsx';
import SaleModal from './SaleModal';
import styles from './Dashboard.module.css';

function Stocks({ stocks }) {
  const [saleOpen, setSaleOpen] = useState(false);
  const [stock, setStock] = useState(undefined);

  const roundNumber = (num) => Math.round((num + Number.EPSILON) * 100) / 100;

  const openSaleModal = (st) => {
    setStock(st);
    setSaleOpen(true);
  };

  return (
    <div style={{ minHeight: '200px' }}>
      <Title>Your Stocks</Title>
      <Table size="small">
        <TableHead>
          <TableRow>
            <TableCell>Ticker</TableCell>
            <TableCell>Quantity</TableCell>
            <TableCell>Purchase Total</TableCell>
            <TableCell>Current Price</TableCell>
            <TableCell align="right">Current Total</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {stocks.map((row) => (
            <TableRow key={row.ticker}>
              <TableCell>
                <Link href="#" onClick={() => openSaleModal(row)}>{row.ticker}</Link>
              </TableCell>
              <TableCell>{row.quantity || '----'}</TableCell>
              <TableCell>
                $
                {row.purchase_total || '----'}
              </TableCell>
              <TableCell>
                $
                {row.current_price.toLocaleString() || '----'}
              </TableCell>
              <TableCell
                align="right"
                className={
                    row.current_balance >= row.purchase_total
                      ? styles.positive
                      : styles.negative
                    }
              >
                $
                {roundNumber(row.current_balance).toLocaleString() || '----'}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      {saleOpen && stock && (
      <SaleModal setSaleOpen={setSaleOpen} stock={stock} />
      )}
    </div>
  );
}

export default Stocks;
