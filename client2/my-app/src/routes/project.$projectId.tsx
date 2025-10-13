import ProjectInfo from "@/components/project/ProjectInfo"
import { Stack, Container } from "@chakra-ui/react"
import { createFileRoute } from "@tanstack/react-router"

export const Route = createFileRoute('/project/$projectId')({
  component: Project,
})

function Project() {
  const { projectId } = Route.useParams()
  return (
    <Stack h="100vh">
      <Container>
        <ProjectInfo ProjectId={projectId}/>
      </Container>
      {/* <Outlet /> */}
    </Stack>
  )
}