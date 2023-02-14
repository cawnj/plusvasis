import base64js from 'base64-js';
import { TextDecoderLite, TextEncoderLite } from 'text-encoder-lite';

export function b64encode(string: string) {
    const encoded = new TextEncoderLite('utf-8').encode(string);
    return base64js.fromByteArray(encoded);
}
export function b64decode(string: string) {
    const uint8array = base64js.toByteArray(string);
    return new TextDecoderLite('utf-8').decode(uint8array);
}
