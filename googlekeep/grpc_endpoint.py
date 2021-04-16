import grpc
from concurrent import futures
import time
import service_pb2_grpc as pb2_grpc
import service_pb2 as pb2
import gkeepapi

def getFromKeep(email,token,name):
        keep = gkeepapi.Keep()
        success = keep.login(email, token)

        gnotes = keep.find(query=name)
        for note in gnotes:
            gnote = note
        # noda = nodes['177b5438cd8.82b17d320d61d816']

        keep.sync()
        return gnote.text.split('\n')

class KeepService(pb2_grpc.KeepServicer):

    def __init__(self, *args, **kwargs):
        pass

    def GetWords(self, request, context):
        # get the string from the incoming request
        email = request.email
        name = request.name
        token = request.token
        wordList = getFromKeep(email, token, name)
        result = {'word': wordList}

        return pb2.SearchResult(**result)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_KeepServicer_to_server(KeepService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()