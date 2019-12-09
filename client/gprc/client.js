const grpc = require("grpc");
const PROTO_PATH = "../../pb/server.proto";
const NoteService = grpc.load(PROTO_PATH).NoteService;

const client = new NoteService(
  "localhost:50051",
  grpc.credentials.createInsecure()
);
module.exports = client;
