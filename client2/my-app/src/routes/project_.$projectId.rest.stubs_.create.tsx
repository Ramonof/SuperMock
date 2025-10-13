import RestStubForm from "@/components/stub/rest/RestStubForm"
import { Stack, Container } from "@chakra-ui/react"
import { createFileRoute } from "@tanstack/react-router"

export const Route = createFileRoute('/project_/$projectId/rest/stubs_/create')({
  component: RestStubs,
})

function RestStubs() {
  const { projectId } = Route.useParams()
  return (
    <Stack h="100vh">
      <Container>
        create form
        <RestStubForm ProjectId={projectId}/>
      </Container>
    </Stack>
  )
}