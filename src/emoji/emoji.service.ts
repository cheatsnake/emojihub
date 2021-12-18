import { Injectable } from "@nestjs/common";
import { InjectConnection, InjectModel } from "@nestjs/mongoose";
import { Collection, connection, Connection, Model, Mongoose } from "mongoose";
import { Emoji, EmojiDocument } from "./schemas/emoji.schema";

@Injectable()
export class EmojiService {
    constructor(
        @InjectModel(Emoji.name) private emojiModel: Model<EmojiDocument>,
        @InjectConnection() private connection: Connection
    ) {}

    async getAll() {
        return connection.collections;
    }
}
