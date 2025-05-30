import {AxiosResponse} from 'axios';
import {CloudResponse} from "@/app/api/models/response/CloudResponse";
import {cloudApi} from '@/app/api/http/cloud';

export default class CloudService {
    static async getAllCloud(type: string): Promise<AxiosResponse<CloudResponse>> {
        return await cloudApi.get<CloudResponse>(`/files/all?type=${type}`);
    }

    static async deleteFile(type: string, obj_id: string) {
        return await cloudApi.delete(`files/one?id=${obj_id}&type=${type}`);
    }
}