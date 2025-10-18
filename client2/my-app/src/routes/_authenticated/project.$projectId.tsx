import ProjectName from "@/components/generic/ProjectName"
import ProjectInfo from "@/components/project/ProjectInfo"
import { Stack, Container } from "@chakra-ui/react"
import { createFileRoute } from "@tanstack/react-router"

export const Route = createFileRoute('/_authenticated/project/$projectId')({
  component: Project,
})

function Project() {
  const { projectId } = Route.useParams()
  return (
    <Stack h="100vh">
      <ProjectName ProjectId={projectId}/>
      <Container>
        <ProjectInfo ProjectId={projectId}/>
      </Container>
      {/* <Outlet /> */}
    </Stack>
  )
}