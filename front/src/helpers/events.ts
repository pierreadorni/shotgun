import type {Event} from '@/types/event'

export function getEvent(id: string): Promise<Event> {
    return new Promise((resolve, reject) => {
        fetch(`http://localhost:3000/api/events/${id}`, {
            credentials: 'include',
        })
            .then((response) => {
                if (response.ok) {
                    response.json().then((eventJson) => {
                        const event: Event = eventJson
                        event.CreatedAt = new Date(event.CreatedAt)
                        event.UpdatedAt = new Date(event.UpdatedAt)
                        event.DeletedAt = new Date(event.DeletedAt)
                        resolve(event)
                    })
                } else {
                    reject(response)
                }
            })
            .catch((error) => {
                reject(error)
            })
    })
}