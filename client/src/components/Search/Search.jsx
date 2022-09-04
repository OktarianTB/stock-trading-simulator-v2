import React, { useState } from 'react';
import {
  TextField, Container,
} from '@mui/material/';
import Autocomplete from '@mui/material/Autocomplete';
import StockCard from './StockCard';
import config from '../../config/Config';

function Search() {
  const [value, setValue] = useState(null);
  const [currentStock, setCurrentStock] = useState(null);

  const onSearchChange = (event, newValue) => {
    setValue(newValue);
    if (newValue) {
      setCurrentStock(newValue);
    } else {
      setCurrentStock(null);
    }
  };

  return (
    <Container>
      <Autocomplete
        value={value}
        onChange={onSearchChange}
        selectOnFocus
        clearOnBlur
        handleHomeEndKeys
        id="stock-search-bar"
        options={config.stocks.sort((a, b) => a.name.localeCompare(b.name))}
        getOptionLabel={(option) => option.name}
        style={{
          maxWidth: '700px',
          margin: '30px auto',
          marginBottom: '60px',
        }}
        renderInput={(params) => (
          <TextField
            {...params}
            label="Search for a stock"
            variant="outlined"
          />
        )}
      />
      {currentStock && (
        <StockCard
          currentStock={currentStock}
        />
      )}
      <br />
      <br />
      <br />
    </Container>
  );
}

export default Search;
