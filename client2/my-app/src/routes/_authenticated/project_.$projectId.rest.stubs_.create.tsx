import ProjectName from "@/components/generic/ProjectName"
import RestStubForm from "@/components/stub/rest/RestStubForm"
import { Stack, Container } from "@chakra-ui/react"
import { createFileRoute } from "@tanstack/react-router"

export const Route = createFileRoute('/_authenticated/project_/$projectId/rest/stubs_/create')({
  component: RestStubs,
})

function RestStubs() {
  const { projectId } = Route.useParams()
  return (
    <Stack h="100vh">
      <ProjectName ProjectId={projectId}/>
      <Container>
        create form
        <RestStubForm ProjectId={projectId}/>
      </Container>
    </Stack>
  )
}