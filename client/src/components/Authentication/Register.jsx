import React, { useState } from 'react';
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
import config from '../../config/Config';
import styles from './Auth.module.css';

function Register() {
  const navigate = useNavigate();

  const [username, setUsername] = useState('');
  const [usernameError, setUsernameError] = useState('');
  const [password, setPassword] = useState('');
  const [passwordError, setPasswordError] = useState('');

  const onChangeUsername = (e) => {
    const newUsername = e.target.value;
    setUsername(newUsername);

    if (newUsername.length < 4 || newUsername.length > 20) {
      setUsernameError('Username must be between 4 and 20 characters.');
    } else {
      setUsernameError('');
    }
  };

  const onChangePassword = (e) => {
    const newPassword = e.target.value;
    setPassword(newPassword);

    if (newPassword.length < 6) {
      setPasswordError('Password must be at least 6 characters.');
    } else {
      setPasswordError('');
    }
  };

  const onSubmit = async (e) => {
    e.preventDefault();
    if (!usernameError && !passwordError) {
      const body = { username, password };
      const url = `${config.base_url}/api/v1/users`;

      await Axios.post(url, body).then(() => {
        navigate('/login');
      }).catch((err) => {
        setUsernameError(err.response.data.error);
        setPasswordError(err.response.data.error);
      });
    }
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
                Register
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
                  error={usernameError.length > 0}
                  helperText={usernameError}
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
                  error={passwordError.length > 0}
                  helperText={passwordError}
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
                    Register
                  </Button>
                </Box>
              </form>
              <Grid container justify="center">
                <Grid item>
                  <Link href="/login" variant="body2">
                    Already have an account?
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

export default Register;
