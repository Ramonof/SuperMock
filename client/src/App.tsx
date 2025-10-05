import { Container, Stack } from '@chakra-ui/react'
import Navbar from './components/navbar'
import ProjectForm from './components/ProjectForm';
import ProjectList from './components/ProjectList';
import { Route, Routes } from 'react-router';
import Projects from './pages/Projects';
import Main from './pages/Main';
import SideBar from './components/SideBar';

export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:8000/api/v1" : "/api";

function App() {

  return (
    <Stack h="100vh">
      <Navbar />
      <SideBar />
      <Routes>
        <Route path='/' element={<Main />} />
        <Route path='/projects' element={<Projects />} />
      </Routes>
      {/* <Container>
        <ProjectForm />
        <ProjectList />
      </Container> */}
    </Stack>
  )
}

export default App
