export interface Abslist {
  AARTIST?: string;
  ALBUM?: string;
  ALBUMID?: string;
  ALIAS?: string;
  ARTIST?: string;
  ARTISTID?: string;
  CanSetRing?: string;
  CanSetRingback?: string;
  DC_TARGETID?: string;
  DC_TARGETTYPE?: string;
  DURATION?: string;
  FARTIST?: string;
  FORMAT?: string;
  FSONGNAME?: string;
  KMARK?: string;
  MINFO?: string;
  MUSICRID?: string;
  MVFLAG?: string;
  MVPIC?: string;
  MVQUALITY?: string;
  NAME?: string;
  NEW?: string;
  N_MINFO?: string;
  ONLINE?: string;
  PAY?: string;
  PROVIDER?: string;
  SONGNAME?: string;
  SUBLIST?: Sublist[];
  SUBTITLE?: string;
  TAG?: string;
  ad_subtype?: string;
  ad_type?: string;
  allartistid?: string;
  audiobookpayinfo?: Audiobookpayinfo;
  barrage?: string;
  cache_status?: string;
  content_type?: string;
  fpay?: string;
  hts_MVPIC?: string;
  info?: string;
  iot_info?: string;
  isdownload?: string;
  isshowtype?: string;
  isstar?: string;
  mvpayinfo?: Mvpayinfo;
  nationid?: string;
  opay?: string;
  originalsongtype?: string;
  overseas_copyright?: string;
  overseas_pay?: string;
  payInfo?: PayInfo;
  react_type?: string;
  spPrivilege?: string;
  subsStrategy?: string;
  subsText?: string;
  terminal?: string;
  tme_musician_adtype?: string;
  tpay?: string;
  web_albumpic_short?: string;
  web_artistpic_short?: string;
  web_timingonline?: string;
}

interface Audiobookpayinfo {
  download: string;
  play: string;
}

interface FeeType {
  album: string;
  bookvip: string;
  song: string;
  vip: string;
}

interface Mvpayinfo {
  download: string;
  play: string;
  vid: string;
}

interface PayInfo {
  cannotDownload: string;
  cannotOnlinePlay: string;
  download: string;
  feeType: FeeType;
  limitfree: string;
  listen_fragment: string;
  local_encrypt: string;
  ndown: string;
  nplay: string;
  overseas_ndown: string;
  overseas_nplay: string;
  paytagindex: Paytagindex;
  play: string;
  refrain_end: string;
  refrain_start: string;
  tips_intercept: string;
}

interface Paytagindex {
  AR501: number;
  DB: number;
  F: number;
  H: number;
  HR: number;
  L: number;
  S: number;
  ZP: number;
  ZPGA201: number;
  ZPGA501: number;
  ZPLY: number;
}

export interface SearchResult {
  ARTISTPIC: string;
  HIT: string;
  HITMODE: string;
  HIT_BUT_OFFLINE: string;
  MSHOW: string;
  NEW: string;
  PN: string;
  RN: string;
  SHOW: string;
  TOTAL: string;
  UK: string;
  abslist: Abslist[];
  searchgroup: string;
}

interface Sublist {
  AARTIST: string;
  ALBUM: string;
  ALBUMID: string;
  ALIAS: string;
  ARTIST: string;
  ARTISTID: string;
  CanSetRing: string;
  CanSetRingback: string;
  DC_TARGETID: string;
  DC_TARGETTYPE: string;
  DURATION: string;
  FARTIST: string;
  FORMAT: string;
  FSONGNAME: string;
  KMARK: string;
  MINFO: string;
  MUSICRID: string;
  MVFLAG: string;
  MVPIC: string;
  MVQUALITY: string;
  NAME: string;
  NEW: string;
  N_MINFO: string;
  ONLINE: string;
  PAY: string;
  PROVIDER: string;
  SONGNAME: string;
  SUBTITLE: string;
  TAG: string;
  ad_subtype: string;
  ad_type: string;
  allartistid: string;
  audiobookpayinfo: Audiobookpayinfo;
  barrage: string;
  cache_status: string;
  content_type: string;
  fpay: string;
  info: string;
  iot_info: string;
  isdownload: string;
  isshowtype: string;
  isstar: string;
  mvpayinfo: Mvpayinfo;
  nationid: string;
  opay: string;
  originalsongtype: string;
  overseas_copyright: string;
  overseas_pay: string;
  payInfo: PayInfo;
  react_type: string;
  spPrivilege: string;
  subsStrategy: string;
  subsText: string;
  terminal: string;
  tme_musician_adtype: string;
  tpay: string;
  web_albumpic_short: string;
  web_artistpic_short: string;
  web_timingonline: string;
}
