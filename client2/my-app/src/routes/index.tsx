import { Flex, useColorModeValue, Stack, Heading, Container, Text, VStack, Box, HStack } from '@chakra-ui/react'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
  component: Index,
})

function Index() {
  return (
        <Stack spacing={10}>
          <Container maxWidth='120ch'>
            <Heading size='2xl' textAlign={'center'}>
              Совершенная платформа симуляции API
            </Heading>
          </Container>
          <Container maxWidth='120ch'>
            <Text fontSize='lg' textAlign={'center'}>
              Запускается везде где запускается ваш код. AI-ready с нативной поддержкой MCP.
            </Text>
          </Container>
          <HStack>
            <Container maxWidth='120ch'>
              <Box bg={useColorModeValue("gray.400", "gray.700")}>
                <Heading size='xl' textAlign={'center'}>
                  Получите
                </Heading>
                <Text>
                  - Тысячи часов в скорости
                </Text>
                <Text>
                  - Более быстрый запуск разработки
                </Text>
                <Text>
                  - Более быстрый запуск разработки
                </Text>
                <Text>
                  - Более надежная скорость параллельной разработки
                </Text>
              </Box>
            </Container>
            <Container maxWidth='120ch'>
              <Box bg={useColorModeValue("gray.400", "gray.700")}>
                <Heading size='xl' textAlign={'center'}>
                  Сократите
                </Heading>
                <Text>
                  - Тысячи часов в скорости
                </Text>
                <Text>
                  - Более быстрый запуск разработки
                </Text>
                <Text>
                  - Более быстрый запуск разработки
                </Text>
                <Text>
                  - Более надежная скорость параллельной разработки
                </Text>
              </Box>
            </Container>
          </HStack>
        </Stack>
  )
}