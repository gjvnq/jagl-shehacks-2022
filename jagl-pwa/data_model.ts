
interface Alter {
    title: string;
    desc: string;
    fronting: boolean;
}

interface Light {
    title: string;
    desc: string;
    color: number;
}

class Person {
    ulid: string = ''
    name: string = ''
    timezone: string = ''
    notices: Array<string> = []
    alters: Array<Alter> = []
    lights: Array<Light> = []
}

export { Alter, Person, Light }
