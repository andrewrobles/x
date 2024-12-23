import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Button, Container, Fab } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import HabitList from '../components/HabitList';
import NewHabitDialog from '../components/NewHabitDialog';
import api from '../api';

function Main() {
  const [habits, setHabits] = useState([]);
  const [error, setError] = useState(null);
  const [openNewHabit, setOpenNewHabit] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchHabits = async () => {
      try {
        const response = await api.getHabits();
        setHabits(response.data);
      } catch (err) {
        if (err.response && err.response.status === 401) {
          navigate('/auth');
        } else {
          setError(err.message);
        }
      }
    };

    fetchHabits();
  }, [navigate]);
  
  const refreshHabits = async () => {
    try {
      const response = await api.getHabits();
      setHabits(response.data);
    } catch (err) {
      setError(err.message);
    }
  };

  const handleSignOut = () => {
    // Clear authentication tokens or perform sign-out logic here
    navigate('/auth'); // Redirect to the authentication page
    localStorage.removeItem('userPassword');
  };

  const handleOpenNewHabit = () => {
    setOpenNewHabit(true);
  };

  const handleCloseNewHabit = () => {
    setOpenNewHabit(false);
  };

  return (
    <div>
      {/* Sign Out Button */}

      {/* Main Content */}
      {error && <p className="error">Error: {error}</p>}
      <Container maxWidth="sm">
        <div style={{ display: 'flex', justifyContent: 'flex-end', padding: '16px' }}>
          <Button color="inherit" onClick={handleSignOut}>
            Sign Out
          </Button>
        </div>
        <HabitList habits={habits} onDelete={refreshHabits} />
        <Fab 
          color="primary" 
          aria-label="add"
          onClick={handleOpenNewHabit}
          sx={{
            position: 'fixed',
            bottom: 16,
            left: '50%',
            transform: 'translateX(-50%)',
            backgroundColor: 'white',
            color: 'primary.main',
            '&:hover': {
              backgroundColor: '#f5f5f5'
            }
          }}
        >
          <AddIcon />
        </Fab>
        <NewHabitDialog
          open={openNewHabit}
          onClose={handleCloseNewHabit}
          onSave={refreshHabits}
        />
      </Container>
    </div>
  );
}

export default Main;