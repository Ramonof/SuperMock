import { Stack, Container } from "@chakra-ui/react"
import Navbar from "../components/navbar"
import ProjectForm from "../components/ProjectForm"
import ProjectList from "../components/ProjectList"

const Projects = () => {
    return (
    <Stack h="100vh">
      <Container>
        <ProjectForm />
        <ProjectList />
      </Container>
    </Stack>
  )
}

export default Projects;