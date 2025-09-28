import { Container, Stack } from '@chakra-ui/react'
import Navbar from './components/navbar'
import ProjectForm from './components/ProjectForm';
import ProjectList from './components/ProjectList';

export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:8000/api/v1" : "/api";

function App() {

  return (
    <Stack h="100vh">
      <Navbar />
      <Container>
        <ProjectForm />
        <ProjectList />
      </Container>
    </Stack>
  )
}

export default App
