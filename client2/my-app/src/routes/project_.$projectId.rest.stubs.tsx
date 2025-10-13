import RestStubList from "@/components/stub/rest/RestStubList"
import StubsInfo from "@/components/stub/rest/StubsInfo"
import { Stack, Container } from "@chakra-ui/react"
import { createFileRoute } from "@tanstack/react-router"

export const Route = createFileRoute('/project_/$projectId/rest/stubs')({
  component: RestStubs,
})

function RestStubs() {
  const { projectId } = Route.useParams()
  return (
    <Stack h="100vh">
      <Container>
        <StubsInfo ProjectId={projectId}/>
        <RestStubList ProjectId={projectId}/>
      </Container>
    </Stack>
  )
}