import { Injectable } from "@nestjs/common";
import { InjectModel } from "@nestjs/mongoose";
import { Model } from "mongoose";
import { Emoji, EmojiDocument } from "./schemas/emoji.schema";

@Injectable()
export class EmojiService {
    constructor(
        @InjectModel(Emoji.name) private emojiModel: Model<EmojiDocument>
    ) {}

    async getAllEmoji(): Promise<Emoji[]> {
        const emojis = await this.emojiModel
            .find({}, { _id: 0 })
            .sort({ _id: 1 })
            .exec();
        return emojis;
    }

    async getEmojiByCategory(category: string): Promise<Emoji[]> {
        const validCategory: string = category.replace(/_/g, " ");
        const emojis = await this.emojiModel
            .find({ category: validCategory }, { _id: 0 })
            .sort({ _id: 1 })
            .exec();
        return emojis;
    }

    async getEmojiByGroup(group: string): Promise<Emoji[]> {
        const validGroup: string = group.replace(/_/g, " ");
        const emojis = await this.emojiModel
            .find({ group: validGroup }, { _id: 0 })
            .sort({ _id: 1 })
            .exec();
        return emojis;
    }
}
