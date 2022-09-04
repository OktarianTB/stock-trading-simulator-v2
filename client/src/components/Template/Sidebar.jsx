import React from 'react';
import {
  Divider,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
} from '@mui/material';
import DashboardIcon from '@mui/icons-material/Dashboard';
import SearchIcon from '@mui/icons-material/Search';
import ExitToAppIcon from '@mui/icons-material/ExitToApp';

function Sidebar({ currentPage, setCurrentPage, logout }) {
  const onDashboardButtonClick = (e) => {
    e.preventDefault();
    setCurrentPage('dashboard');
  };

  const onSearchButtonClick = (e) => {
    e.preventDefault();
    setCurrentPage('search');
  };

  return (
    <div>
      <List>
        <ListItem
          button
          selected={currentPage === 'dashboard'}
          onClick={onDashboardButtonClick}
        >
          <ListItemIcon>
            <DashboardIcon />
          </ListItemIcon>
          <ListItemText primary="Dashboard" />
        </ListItem>
        <ListItem
          button
          selected={currentPage === 'search'}
          onClick={onSearchButtonClick}
        >
          <ListItemIcon>
            <SearchIcon />
          </ListItemIcon>
          <ListItemText primary="Search" />
        </ListItem>
      </List>
      <Divider />
      <List>
        <ListItem button onClick={logout}>
          <ListItemIcon>
            <ExitToAppIcon />
          </ListItemIcon>
          <ListItemText primary="Log Out" />
        </ListItem>
      </List>
    </div>
  );
}

export default Sidebar;
