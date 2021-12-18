import { Prop, Schema, SchemaFactory } from "@nestjs/mongoose";
import { Document } from "mongoose";

export type EmojiDocument = Emoji & Document;

@Schema()
export class Emoji {
    @Prop()
    name: string;

    @Prop()
    category: string;

    @Prop()
    group: string;

    @Prop([String])
    htmlCode: string[];

    @Prop([String])
    unicode: string[];
}

export const EmojiSchema = SchemaFactory.createForClass(Emoji);
