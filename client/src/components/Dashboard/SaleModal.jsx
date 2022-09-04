import React, { useState, useContext } from 'react';
import {
  Typography,
  IconButton,
  Box,
  Button,
  TextField,
  Container,
  Grid,
  Card,
  CardHeader,
  CardContent,
} from '@mui/material';
import { motion } from 'framer-motion';
import CloseIcon from '@mui/icons-material/Close';
import Axios from 'axios';
import styles from '../Template/PageTemplate.module.css';
import UserContext from '../../context/UserContext';
import config from '../../config/Config';

function SaleModal({ setSaleOpen, stock }) {
  return (
    <motion.div
      className={styles.backdrop}
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      id="backdrop"
    >
      <Container>
        <motion.div animate={{ opacity: 1, y: -20 }}>
          <SaleModalContent setSaleOpen={setSaleOpen} stock={stock} />
        </motion.div>
      </Container>
    </motion.div>
  );
}

function SaleModalContent({ setSaleOpen, stock }) {
  const { userData } = useContext(UserContext);
  const [quantity, setQuantity] = useState(1);

  const handleQuantityChange = (e) => {
    if (!Number.isNaN(e.target.value) && Number(e.target.value) <= stock.quantity
    && Number(e.target.value) >= 1) {
      setQuantity(Number(e.target.value));
    }
  };

  const handleClick = () => {
    setSaleOpen(false);
  };

  const sellStock = async (e) => {
    e.preventDefault();

    const headers = { Authorization: `Bearer ${userData.accessToken}` };
    const body = { ticker: stock.ticker, quantity };

    await Axios.post(
      `${config.base_url}/api/v1/transactions/sell`,
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
        <Card>
          <CardHeader
            action={(
              <IconButton aria-label="Close" onClick={handleClick}>
                <CloseIcon />
              </IconButton>
            )}
          />
          <CardContent>
            <Typography component="h1" variant="h6" align="center">
              Sell
            </Typography>
            <form className={styles.form} onSubmit={(e) => e.preventDefault()}>
              <TextField
                variant="outlined"
                margin="normal"
                fullWidth
                disabled
                id="name"
                label="Name"
                name="Name"
                autoComplete="Name"
                value={stock.ticker}
              />
              <TextField
                variant="outlined"
                margin="normal"
                fullWidth
                disabled
                id="price"
                label="Price"
                name="price"
                autoComplete="price"
                value={stock.current_price}
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
                value={quantity}
                type="number"
                onChange={handleQuantityChange}
              />
            </form>
            <br />
            <Box display="flex" justifyContent="center">
              <Button
                type="submit"
                variant="contained"
                color="primary"
                className={styles.confirm}
                onClick={sellStock}
              >
                Confirm
              </Button>
            </Box>

            <br />
            <br />
          </CardContent>
        </Card>
      </Box>
    </Grid>
  );
}

export default SaleModal;
