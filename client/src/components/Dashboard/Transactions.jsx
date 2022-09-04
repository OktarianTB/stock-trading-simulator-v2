import React from 'react';
import {
  Table, TableBody, TableCell, TableHead, TableRow,
} from '@mui/material';
import Title from '../Template/Title.jsx';

function Transactions({ transactions }) {
  const roundNumber = (num) => Math.round((num + Number.EPSILON) * 100) / 100;

  return (
    <div style={{ minHeight: '200px' }}>
      <Title>Past Transactions</Title>
      <Table size="small">
        <TableHead>
          <TableRow>
            <TableCell>Type</TableCell>
            <TableCell align="right">Ticker</TableCell>
            <TableCell align="right">Price</TableCell>
            <TableCell align="right">Quantity</TableCell>
            <TableCell align="right">Purchased At</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {transactions.map((row, i) => (
            <TableRow key={i}>
              <TableCell>{row.quantity > 0 ? 'BUY' : 'SELL'}</TableCell>
              <TableCell align="right">
                {row.ticker}
              </TableCell>
              <TableCell align="right">
                $
                {roundNumber(row.price).toLocaleString()}
              </TableCell>
              <TableCell align="right">
                {row.quantity > 0 && '+'}
                {row.quantity}
              </TableCell>
              <TableCell align="right">
                {new Date(Date.parse(row.purchased_at)).toLocaleDateString('en-UK')}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}

export default Transactions;
