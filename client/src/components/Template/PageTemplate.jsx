import React, { useContext, useState } from 'react';
import { styled } from '@mui/material/styles';
import { useNavigate } from 'react-router-dom';
import {
  Drawer,
  CssBaseline,
  Box,
  AppBar,
  Toolbar,
  Typography,
  Divider,
  IconButton,
} from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import styles from './PageTemplate.module.css';
import UserContext from '../../context/UserContext';
import Sidebar from './Sidebar';
import Dashboard from '../Dashboard/Dashboard';
import Search from '../Search/Search';

const drawerWidth = 240;

const Main = styled('main', { shouldForwardProp: (prop) => prop !== 'open' })(
  ({ theme, open }) => ({
    flexGrow: 1,
    padding: theme.spacing(3),
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    marginLeft: `-${drawerWidth}px`,
    ...(open && {
      transition: theme.transitions.create('margin', {
        easing: theme.transitions.easing.easeOut,
        duration: theme.transitions.duration.enteringScreen,
      }),
      marginLeft: 0,
    }),
  }),
);

const Menu = styled(AppBar, {
  shouldForwardProp: (prop) => prop !== 'open',
})(({ theme, open }) => ({
  transition: theme.transitions.create(['margin', 'width'], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: `${drawerWidth}px`,
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));

const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: 'flex-end',
}));

function PageTemplate() {
  const navigate = useNavigate();
  const { userData, setUserData } = useContext(UserContext);
  const [open, setOpen] = useState(true);
  const [currentPage, setCurrentPage] = useState('dashboard');

  if (!userData.user) {
    navigate('/login');
  }

  const logout = () => {
    setUserData({
      token: undefined,
      user: undefined,
    });
    localStorage.setItem('refresh-token', '');
    navigate('/login');
  };

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  return (
    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      <Menu position="fixed" open={open}>
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            onClick={handleDrawerOpen}
            edge="start"
            sx={{ mr: 2, ...(open && { display: 'none' }) }}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" noWrap component="div" className={styles.title}>
            {currentPage === 'dashboard' && 'Dashboard'}
            {currentPage === 'search' && 'Search'}
          </Typography>
          <Typography color="inherit">
            Hello,
            {' '}
            {userData.user.username
              ? userData.user.username.charAt(0).toUpperCase()
                  + userData.user.username.slice(1)
              : ''}
          </Typography>
        </Toolbar>
      </Menu>
      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
          },
        }}
        variant="persistent"
        anchor="left"
        open={open}
      >
        <DrawerHeader>
          <IconButton onClick={handleDrawerClose}>
            <ChevronLeftIcon />
          </IconButton>
        </DrawerHeader>
        <Divider />
        <Sidebar currentPage={currentPage} setCurrentPage={setCurrentPage} logout={logout} />
      </Drawer>
      <Main open={open}>
        <DrawerHeader />
        {currentPage === 'dashboard' && (
          <Dashboard />
        )}
        {currentPage === 'search' && (
          <Search />
        )}
      </Main>
    </Box>
  );
}

export default PageTemplate;
