import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Axios from 'axios';
import {
  Login, Register, NotFound, PageTemplate,
} from './components';
import UserContext from './context/UserContext';
import config from './config/Config';

function App() {
  const [userData, setUserData] = useState({
    accessToken: undefined,
    user: undefined,
  });

  useEffect(() => {
    const checkLoggedIn = async () => {
      const refreshToken = localStorage.getItem('refresh-token');
      if (refreshToken == null || refreshToken.length === 0) {
        localStorage.setItem('refresh-token', '');
        setUserData({ token: undefined, user: undefined });
        return;
      }

      const accessToken = await Axios.post(
        `${config.base_url}/api/v1/tokens/renew_access`,
        { refresh_token: refreshToken },
        null,
      );

      if (accessToken?.data?.access_token) {
        const headers = { Authorization: `Bearer ${accessToken.data.access_token}` };

        await Axios.get(
          `${config.base_url}/api/v1/users`,
          { headers },
        ).then((res) => {
          setUserData({
            accessToken: accessToken.data.access_token,
            user: res.data,
          });
        });
      }
    };

    checkLoggedIn();
  }, []);

  return (
    <Router>
      <UserContext.Provider value={{ userData, setUserData }}>
        <Routes>
          {userData.user ? (
            <Route path="/" element={<PageTemplate />} />
          ) : (
            <Route path="/" element={<Login />} />
          )}
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </UserContext.Provider>
    </Router>
  );
}

export default App;
