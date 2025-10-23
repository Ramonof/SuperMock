import { BASE_URL } from "@/main";
import { Flex, Input, Button, Spinner, Stack, Text, HStack, Select } from "@chakra-ui/react";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import { useEffect, useRef, useState } from "react";
import { IoMdAdd } from "react-icons/io";
import type { RestStub } from "./RestStubList";
import apiClient from "@/utils/request";
import {EditorView, basicSetup} from "codemirror"
import {json, jsonParseLinter} from "@codemirror/lang-json"
import {oneDark} from "@codemirror/theme-one-dark"
import CodeMirror from '@uiw/react-codemirror';
import React from "react";
import { linter } from "@codemirror/lint";

const RestStubFullForm = ({ ProjectId, StubId, restStubs, setRestStubs }: { ProjectId: string, StubId: string, restStubs: RestStub[], setRestStubs: React.Dispatch<React.SetStateAction<RestStub[]>> }) => {

    const [value, setValue] = React.useState("console.log('hello world!');");
    const onChange = React.useCallback((val: React.SetStateAction<string>, viewUpdate: any) => {
        console.log('val:', val);
        setValue(val);
        setNewRestStubResponseBody(val)
    }, []);

    const [newRestStubName, setNewRestStubName] = useState("");
    const [newRestStubPath, setNewRestStubPath] = useState("");
    const [newRestStubMethod, setNewRestStubMethod] = useState("");
    const [newRestStubResponseBody, setNewRestStubResponseBody] = useState("");

    const navigate = useNavigate();

    const queryClient = useQueryClient();

    const { data: stubData, isLoading } = useQuery<RestStub>({
		queryKey: ["getStubById" + StubId],
		queryFn: async () => {
			try {
				const response = await apiClient({
					url: BASE_URL + "/projects/" + ProjectId + "/stub/" + StubId, 
					method: 'get'
				});
				// const res = response.data
				const data = response.data
                setNewRestStubName(data.name);
                setNewRestStubPath(data.path);
                setNewRestStubMethod(data.method);
                setNewRestStubResponseBody(data.response_body)

				// if (!res.ok) {
				// 	throw new Error(data.error || "Something went wrong");
				// }
				return data || null;
			} catch (error) {
				console.log(error);
			}
		},
	});

    const { mutate: updateRestStub, isPending: isCreating } = useMutation({
        mutationKey: ["updateRestStub"],
        mutationFn: async (e: React.FormEvent) => {
            e.preventDefault();
            try {
                const res = await apiClient({
					url: BASE_URL + "/projects/" + ProjectId + "/stub/" + StubId, 
					method: 'PUT',
                    headers: {
                        "Content-Type": "application/json",
                    },
                    data: JSON.stringify({ name: newRestStubName, path: newRestStubPath, method: newRestStubMethod, response_body: newRestStubResponseBody }),
				});

                setRestStubs(restStubs.map(item => (item.id == res.data.id ? res.data : item)));
                
                // setNewRestStubName("");
                // setNewRestStubPath("");
                // setNewRestStubMethod("")
                // setNewRestStubResponseBody("");
                // navigate({
                //     to: '/project/$projectId/rest/stubs/$stubId',
                //     params: { projectId: ProjectId, stubId: StubId },
                // })
                return res;
            } catch (error: any) {
                throw new Error(error);
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ["getRestStubs"] });
        },
        onError: (error: any) => {
            alert(error.message);
        },
    });

    return (
        <form onSubmit={updateRestStub}>
            <Stack gap={2}>
                <Text color={"yellow.100"}  fontSize='xl'>
                    Request
                </Text>
                <Text color={"yellow.100"}>
                    Name
                </Text>
                <Input
                    type='text'
                    value={newRestStubName}
                    onChange={(e) => setNewRestStubName(e.target.value)}
                    // ref={(input) => input && input.focus()}
                    textFillColor={"yellow.100"}
                />
                <Text color={"yellow.100"}>
                    Path
                </Text>
                <HStack>
                    <Select value={newRestStubMethod} onChange={(e) => setNewRestStubMethod(e.target.value)} color={"yellow.100"} width={'fit-content'}>
                        <option value='ANY'>ANY</option>
                        <option value='GET'>GET</option>
                        <option value='POST'>POST</option>
                        <option value='PUT'>PUT</option>
                        <option value='PATCH'>PATCH</option>
                        <option value='DELETE'>DELETE</option>
                        <option value='HEAD'>HEAD</option>
                        <option value='OPTIONS'>OPTIONS</option>
                    </Select>
                    <Input
                        type='text'
                        value={newRestStubPath}
                        onChange={(e) => setNewRestStubPath(e.target.value)}
                        // ref={(input) => input && input.focus()}
                        textFillColor={"yellow.100"}
                    />
                </HStack>
                <Text color={"yellow.100"}  fontSize='xl'>
                    Response
                </Text>
                <CodeMirror value={newRestStubResponseBody} height="200px" extensions={[json(), linter(jsonParseLinter())]} onChange={onChange} />
                <Button
                    mx={2}
                    type='submit'
                    _active={{
                        transform: "scale(.97)",
                    }}
                >
                    {isCreating ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />} Save
                </Button>
            </Stack>
        </form>
    );
};
export default RestStubFullForm;