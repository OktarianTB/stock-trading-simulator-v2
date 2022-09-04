import React, { useState, useContext } from 'react';
import {
  Box,
  Typography,
  TextField,
  CssBaseline,
  Button,
  Card,
  CardContent,
  Grid,
  Link,
} from '@mui/material';
import { useNavigate } from 'react-router-dom';
import Axios from 'axios';
import UserContext from '../../context/UserContext';
import config from '../../config/Config';

import styles from './Auth.module.css';

function Login() {
  const navigate = useNavigate();
  const { setUserData } = useContext(UserContext);

  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [loginError, setLoginError] = useState('');

  const onChangeUsername = (e) => {
    const newUsername = e.target.value;
    setUsername(newUsername);
  };

  const onChangePassword = (e) => {
    const newPassword = e.target.value;
    setPassword(newPassword);
  };

  const onSubmit = async (e) => {
    e.preventDefault();
    const body = { username, password };
    const url = `${config.base_url}/api/v1/users/login`;

    await Axios.post(url, body).then((res) => {
      setLoginError('');
      setUserData({
        accessToken: res.data.access_token,
        user: res.data.user,
      });
      localStorage.setItem('refresh-token', res.data.refresh_token);
      navigate('/');
    }).catch(() => {
      setLoginError('Invalid username or password');
    });
  };

  return (
    <div className={styles.background}>
      <CssBaseline />
      <Grid
        container
        spacing={0}
        direction="column"
        alignItems="center"
        justify="center"
        style={{ minHeight: '100vh' }}
      >
        <Box width="70vh" boxShadow={1} className={styles.box}>
          <Card className={styles.paper}>
            <CardContent>
              <Typography component="h1" variant="h5">
                Login
              </Typography>
              <form className={styles.form} onSubmit={onSubmit}>
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  id="username"
                  label="Username"
                  name="username"
                  autoComplete="username"
                  error={loginError.length > 0}
                  helperText={loginError}
                  value={username}
                  onChange={onChangeUsername}
                />
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  name="password"
                  label="Password"
                  type="password"
                  id="password"
                  autoComplete="current-password"
                  error={loginError.length > 0}
                  helperText={loginError}
                  value={password}
                  onChange={onChangePassword}
                />
                <Box display="flex" justifyContent="center">
                  <Button
                    type="submit"
                    variant="contained"
                    color="primary"
                    className={styles.submit}
                  >
                    Login
                  </Button>
                </Box>
              </form>
              <Grid container justify="center">
                <Grid item>
                  <Link href="/register" variant="body2">
                    Need an account?
                  </Link>
                </Grid>
              </Grid>
            </CardContent>
          </Card>
        </Box>
      </Grid>
    </div>
  );
}

export default Login;
