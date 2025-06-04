import { RecordDataResponse } from "@/app/types/apiData";
import { ContactData, ContactLink } from "@/app/types";

export function extractContactData(record: RecordDataResponse): ContactData{
   
    
    const emails = record.person.contact.email;
    const primaryEmail = emails? emails[0]: ""; //pega o primeiro email
    
    //array de urls
    const rawsUrls = record.person.contact["researcher-url"] ?? [];

    const links: ContactLink[] = rawsUrls.map((item)=>{
        return{
            name: item["url-name"] || item.url,
            url: item.url,
            icon: ""
        }
    })

    return {
        email: primaryEmail,
        links,
    }
}