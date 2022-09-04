import config from '../config/Config';

const randomInteger = (min, max) => Math.floor(Math.random() * (max - min + 1)) + min;

const getRandomTicker = () => {
  const nb = randomInteger(0, config.stocks.length - 1);
  return config.stocks[nb].ticker;
};

export default getRandomTicker;
