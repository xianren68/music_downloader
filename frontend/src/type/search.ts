export interface Albums {
    id: string;
    name: string;
    type: string;
  }
  
  interface ImgItems {
    imgSizeType: string;
    img: string;
  }
  
  interface MvList {
    id: string;
    copyrightId: string;
    resourceType: string;
    price: string;
    expireDate: string;
    mvPicUrl: MvPicUrl[];
    isInDAlbum: string;
    playNum: string;
    mvType: string;
  }
  
  interface MvPicUrl {
    imgSizeType: string;
    img: string;
    fileId: string;
  }
  
  interface RateFormats {
    resourceType: string;
    formatType: string;
    format: string;
    size: string;
    fileType?: string;
    price: string;
    showTag: string[];
    androidFileType?: string;
    iosFileType?: string;
    iosSize?: string;
    androidSize?: string;
    iosFormat?: string;
    androidFormat?: string;
    iosAccuracyLevel?: string;
    androidAccuracyLevel?: string;
  }
  
  interface RelatedSongs {
    resourceType: string;
    resourceTypeName: string;
    copyrightId: string;
    productId: string;
  }
  
export  interface Result {
    id: string;
    resourceType: string;
    contentId: string;
    copyrightId: string;
    name: string;
    highlightStr: string[];
    singers: Singers[];
    albums: Albums[];
    lyricUrl: string;
    trcUrl: string;
    imgItems: ImgItems[];
    televisionNames: string[];
    tones: Tones[];
    mvList?: MvList[];
    relatedSongs: RelatedSongs[];
    toneControl: string;
    rateFormats: RateFormats[];
    newRateFormats: RateFormats[];
    songType: string;
    isInDAlbum: string;
    copyright: string;
    digitalColumnId: string;
    mrcurl: string;
    songDescs: string;
    songAliasName: string;
    invalidateDate: string;
    isInSalesPeriod: string;
    dalbumId: string;
    isInSideDalbum: string;
    vipType: string;
    chargeAuditions: string;
    scopeOfcopyright: string;
    mvCopyright?: string;
    translateName?: string;
    z3dCode?: Z3DCode;
    tags?: string[];
    movieNames?: string[];
  }
  
  interface RootInterface {
    code: string;
    info: string;
    songResultData: SongResultData;
    tagSongResultData: TagSongResultData;
    bestShowResultToneData: any;
  }
  
export interface Singers {
    id: string;
    name: string;
  }
  
  interface SongResultData {
    totalCount: string;
    correct: any[];
    resultType: string;
    isFromCache: string;
    result: Result[];
    tipStatus: string;
  }
  
  interface TagSongResultData {
    correct: any[];
    result: any[];
  }
  
  interface Tones {
    id: string;
    copyrightId: string;
    price: string;
    expireDate: string;
  }
  
  interface Z3DCode {
    resourceType: string;
    formatType: string;
    price: string;
    androidFileType: string;
    iosFileType: string;
    iosSize: string;
    androidSize: string;
    iosFormat: string;
    androidFormat: string;
    androidFileKey: string;
    iosFileKey: string;
    h5Size: string;
    h5Format: string;
  }

 export interface SearchResult{
    code:string,
    info:string,
    songResultData:SongResultData,
    tagSongResultData:TagSongResultData,
    bestShowResultToneData:any,
  }