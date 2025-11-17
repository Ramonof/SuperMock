import { Box, HStack, Link, Stack, Text } from '@chakra-ui/react'
import { createFileRoute } from '@tanstack/react-router'
import { IoAddCircleOutline } from 'react-icons/io5'
import getColor from "@/utils/color";
import { Link as TanstackLink } from "@tanstack/react-router";

export const Route = createFileRoute(
  '/_authenticated/project_/$projectId/variables',
)({
  component: Variables,
})

function Variables() {
  const { projectId } = Route.useParams()
  return <div>
    <Stack>
      <Box
        maxW='fit-content' borderWidth='1px' borderRadius='lg' overflow='hidden'
            border={"1px"}
            borderColor={"gray.600"}
            p={3}
      >
        <Link as={TanstackLink}
          to={`/project/${projectId}/rest/stubs/create`}
          color={getColor()}
          variant="underline"
          colorPalette="teal"
        >
          <HStack>
            <IoAddCircleOutline />
            <Text>Add Variable</Text>
          </HStack>
        </Link>
      </Box>
      <Text>
        Variables
      </Text>
      <Text>
        res1 : {'{'}"test":"body"{'}'}
      </Text>
    </Stack>
    </div>
}
