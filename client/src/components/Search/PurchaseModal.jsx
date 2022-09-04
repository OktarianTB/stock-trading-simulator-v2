import React, { useState, useContext } from 'react';
import {
  Container,
  Typography,
  Box,
  Button,
  TextField,
  CardContent,
  CardHeader,
  IconButton,
  Grid,
  Card,
} from '@mui/material/';
import CloseIcon from '@mui/icons-material/Close';
import { motion } from 'framer-motion';
import Axios from 'axios';
import styles from './Search.module.css';
import UserContext from '../../context/UserContext';
import config from '../../config/Config';

function PurchaseModal({
  stockMetadata,
  price,
  setSelected,
}) {
  return (
    <motion.div
      className={styles.backdrop}
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      id="backdrop"
    >
      <Container>
        <motion.div animate={{ opacity: 1, y: -20 }}>
          <PurchaseModalContent
            stockMetadata={stockMetadata}
            price={price}
            setSelected={setSelected}
          />
        </motion.div>
      </Container>
    </motion.div>
  );
}

function PurchaseModalContent({
  stockMetadata,
  price,
  setSelected,
}) {
  const [quantity, setQuantity] = useState(1);
  const [total, setTotal] = useState(Number(price));
  const { userData } = useContext(UserContext);

  const handleQuantityChange = (e) => {
    if (!Number.isNaN(e.target.value) && Number(e.target.value) >= 1) {
      if (
        userData.user.balance
          - Number(price) * Number(e.target.value)
        < 0
      ) {
        return;
      }

      setQuantity(Number(e.target.value));
      setTotal(Number(price) * Number(e.target.value));
    }
  };

  const handleClick = () => {
    setSelected(false);
  };

  const purchaseStock = async (e) => {
    e.preventDefault();

    const headers = { Authorization: `Bearer ${userData.accessToken}` };
    const body = { ticker: stockMetadata.ticker, quantity };

    await Axios.post(
      `${config.base_url}/api/v1/transactions/purchase`,
      body,
      { headers },
    ).then(() => {
      window.location.reload();
    });
  };

  return (
    <Grid
      container
      spacing={0}
      direction="column"
      alignItems="center"
      justify="center"
      style={{ minHeight: '100vh', marginTop: '20vh' }}
    >
      <Box width="60vh" boxShadow={1}>
        <Card className={styles.paper}>
          <CardHeader
            action={(
              <IconButton aria-label="Close" onClick={handleClick}>
                <CloseIcon />
              </IconButton>
            )}
          />
          <CardContent>
            <Typography component="h1" variant="h6" align="center">
              Purchase
              {' '}
              {stockMetadata.name}
              {' '}
              Stock
            </Typography>
            <form className={styles.form} onSubmit={(e) => e.preventDefault()}>
              <TextField
                variant="outlined"
                margin="normal"
                fullWidth
                disabled
                id="stock"
                label="Stock Name"
                name="stock"
                autoComplete="stock"
                value={stockMetadata.name}
              />
              <TextField
                variant="outlined"
                margin="normal"
                fullWidth
                disabled
                id="price"
                label="Stock Price"
                name="price"
                autoComplete="price"
                value={Math.round(price * 100) / 100}
              />
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="quantity"
                label="Quantity"
                name="quantity"
                autoComplete="quantity"
                type="number"
                value={quantity}
                onChange={handleQuantityChange}
              />
              <Typography
                variant="body2"
                align="center"
                className={styles.addMargin}
              >
                Total = $
                {(Math.round(total * 100) / 100).toLocaleString()}
              </Typography>
              <Typography variant="body2" align="center">
                Cash balance after purchase:
                {' '}
                {userData
                  ? `$${(userData.user.balance - total).toLocaleString()}`
                  : 'Balance Unavailable'}
              </Typography>
              <Box display="flex" justifyContent="center">
                <Button
                  type="submit"
                  variant="contained"
                  color="primary"
                  className={styles.submit}
                  onClick={purchaseStock}
                >
                  Confirm
                </Button>
              </Box>
            </form>
            <br />
            <br />
          </CardContent>
        </Card>
      </Box>
    </Grid>
  );
}

export default PurchaseModal;
