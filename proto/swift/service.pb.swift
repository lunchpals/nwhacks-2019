// DO NOT EDIT.
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: service.proto
//
// For information on using the generated types, please see the documenation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that your are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

struct OpenNow_Status {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct OpenNow_Empty {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct OpenNow_Coordinates {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var latitude: Double = 0

  var longitude: Double = 0

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct OpenNow_Position {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var clientID: String {
    get {return _storage._clientID}
    set {_uniqueStorage()._clientID = newValue}
  }

  var situation: OpenNow_Context {
    get {return _storage._situation ?? OpenNow_Context()}
    set {_uniqueStorage()._situation = newValue}
  }
  /// Returns true if `situation` has been explicitly set.
  var hasSituation: Bool {return _storage._situation != nil}
  /// Clears the value of `situation`. Subsequent reads from it will return its default value.
  mutating func clearSituation() {_uniqueStorage()._situation = nil}

  var direction: Float {
    get {return _storage._direction}
    set {_uniqueStorage()._direction = newValue}
  }

  var coordinates: OpenNow_Coordinates {
    get {return _storage._coordinates ?? OpenNow_Coordinates()}
    set {_uniqueStorage()._coordinates = newValue}
  }
  /// Returns true if `coordinates` has been explicitly set.
  var hasCoordinates: Bool {return _storage._coordinates != nil}
  /// Clears the value of `coordinates`. Subsequent reads from it will return its default value.
  mutating func clearCoordinates() {_uniqueStorage()._coordinates = nil}

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

struct OpenNow_Context {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var situation: OpenNow_Context.Situation = .unknown

  var unknownFields = SwiftProtobuf.UnknownStorage()

  enum Situation: SwiftProtobuf.Enum {
    typealias RawValue = Int
    case unknown // = 0
    case vehicle // = 1
    case foot // = 2
    case UNRECOGNIZED(Int)

    init() {
      self = .unknown
    }

    init?(rawValue: Int) {
      switch rawValue {
      case 0: self = .unknown
      case 1: self = .vehicle
      case 2: self = .foot
      default: self = .UNRECOGNIZED(rawValue)
      }
    }

    var rawValue: Int {
      switch self {
      case .unknown: return 0
      case .vehicle: return 1
      case .foot: return 2
      case .UNRECOGNIZED(let i): return i
      }
    }

  }

  init() {}
}

#if swift(>=4.2)

extension OpenNow_Context.Situation: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static var allCases: [OpenNow_Context.Situation] = [
    .unknown,
    .vehicle,
    .foot,
  ]
}

#endif  // swift(>=4.2)

struct OpenNow_PointsOfInterest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var interests: [OpenNow_Interest] = []

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct OpenNow_Interest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var interestID: String {
    get {return _storage._interestID}
    set {_uniqueStorage()._interestID = newValue}
  }

  var name: String {
    get {return _storage._name}
    set {_uniqueStorage()._name = newValue}
  }

  var description_p: String {
    get {return _storage._description_p}
    set {_uniqueStorage()._description_p = newValue}
  }

  var photoRef: String {
    get {return _storage._photoRef}
    set {_uniqueStorage()._photoRef = newValue}
  }

  var type: OpenNow_Interest.TypeEnum {
    get {return _storage._type}
    set {_uniqueStorage()._type = newValue}
  }

  var coordinates: OpenNow_Coordinates {
    get {return _storage._coordinates ?? OpenNow_Coordinates()}
    set {_uniqueStorage()._coordinates = newValue}
  }
  /// Returns true if `coordinates` has been explicitly set.
  var hasCoordinates: Bool {return _storage._coordinates != nil}
  /// Clears the value of `coordinates`. Subsequent reads from it will return its default value.
  mutating func clearCoordinates() {_uniqueStorage()._coordinates = nil}

  var unknownFields = SwiftProtobuf.UnknownStorage()

  enum TypeEnum: SwiftProtobuf.Enum {
    typealias RawValue = Int
    case unknown // = 0
    case authority // = 1
    case food // = 2
    case store // = 3
    case lodging // = 4
    case attraction // = 5
    case UNRECOGNIZED(Int)

    init() {
      self = .unknown
    }

    init?(rawValue: Int) {
      switch rawValue {
      case 0: self = .unknown
      case 1: self = .authority
      case 2: self = .food
      case 3: self = .store
      case 4: self = .lodging
      case 5: self = .attraction
      default: self = .UNRECOGNIZED(rawValue)
      }
    }

    var rawValue: Int {
      switch self {
      case .unknown: return 0
      case .authority: return 1
      case .food: return 2
      case .store: return 3
      case .lodging: return 4
      case .attraction: return 5
      case .UNRECOGNIZED(let i): return i
      }
    }

  }

  init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

#if swift(>=4.2)

extension OpenNow_Interest.TypeEnum: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static var allCases: [OpenNow_Interest.TypeEnum] = [
    .unknown,
    .authority,
    .food,
    .store,
    .lodging,
    .attraction,
  ]
}

#endif  // swift(>=4.2)

struct OpenNow_DirectionsReq {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var clientID: String = String()

  var interestID: String = String()

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct OpenNow_DirectionsResp {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var interestID: String = String()

  var legs: [Data] = []

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "open_now"

extension OpenNow_Status: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Status"
  static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: OpenNow_Status, rhs: OpenNow_Status) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension OpenNow_Empty: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Empty"
  static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: OpenNow_Empty, rhs: OpenNow_Empty) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension OpenNow_Coordinates: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Coordinates"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "latitude"),
    2: .same(proto: "longitude"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 1: try decoder.decodeSingularDoubleField(value: &self.latitude)
      case 2: try decoder.decodeSingularDoubleField(value: &self.longitude)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.latitude != 0 {
      try visitor.visitSingularDoubleField(value: self.latitude, fieldNumber: 1)
    }
    if self.longitude != 0 {
      try visitor.visitSingularDoubleField(value: self.longitude, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: OpenNow_Coordinates, rhs: OpenNow_Coordinates) -> Bool {
    if lhs.latitude != rhs.latitude {return false}
    if lhs.longitude != rhs.longitude {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension OpenNow_Position: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Position"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "client_id"),
    2: .same(proto: "situation"),
    3: .same(proto: "direction"),
    4: .same(proto: "coordinates"),
  ]

  fileprivate class _StorageClass {
    var _clientID: String = String()
    var _situation: OpenNow_Context? = nil
    var _direction: Float = 0
    var _coordinates: OpenNow_Coordinates? = nil

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _clientID = source._clientID
      _situation = source._situation
      _direction = source._direction
      _coordinates = source._coordinates
    }
  }

  fileprivate mutating func _uniqueStorage() -> _StorageClass {
    if !isKnownUniquelyReferenced(&_storage) {
      _storage = _StorageClass(copying: _storage)
    }
    return _storage
  }

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    _ = _uniqueStorage()
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      while let fieldNumber = try decoder.nextFieldNumber() {
        switch fieldNumber {
        case 1: try decoder.decodeSingularStringField(value: &_storage._clientID)
        case 2: try decoder.decodeSingularMessageField(value: &_storage._situation)
        case 3: try decoder.decodeSingularFloatField(value: &_storage._direction)
        case 4: try decoder.decodeSingularMessageField(value: &_storage._coordinates)
        default: break
        }
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      if !_storage._clientID.isEmpty {
        try visitor.visitSingularStringField(value: _storage._clientID, fieldNumber: 1)
      }
      if let v = _storage._situation {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
      }
      if _storage._direction != 0 {
        try visitor.visitSingularFloatField(value: _storage._direction, fieldNumber: 3)
      }
      if let v = _storage._coordinates {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 4)
      }
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: OpenNow_Position, rhs: OpenNow_Position) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._clientID != rhs_storage._clientID {return false}
        if _storage._situation != rhs_storage._situation {return false}
        if _storage._direction != rhs_storage._direction {return false}
        if _storage._coordinates != rhs_storage._coordinates {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension OpenNow_Context: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Context"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "situation"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 1: try decoder.decodeSingularEnumField(value: &self.situation)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.situation != .unknown {
      try visitor.visitSingularEnumField(value: self.situation, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: OpenNow_Context, rhs: OpenNow_Context) -> Bool {
    if lhs.situation != rhs.situation {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension OpenNow_Context.Situation: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "UNKNOWN"),
    1: .same(proto: "VEHICLE"),
    2: .same(proto: "FOOT"),
  ]
}

extension OpenNow_PointsOfInterest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".PointsOfInterest"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "interests"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 1: try decoder.decodeRepeatedMessageField(value: &self.interests)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.interests.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.interests, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: OpenNow_PointsOfInterest, rhs: OpenNow_PointsOfInterest) -> Bool {
    if lhs.interests != rhs.interests {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension OpenNow_Interest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Interest"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "interest_id"),
    2: .same(proto: "name"),
    3: .same(proto: "description"),
    4: .standard(proto: "photo_ref"),
    6: .same(proto: "type"),
    7: .same(proto: "coordinates"),
  ]

  fileprivate class _StorageClass {
    var _interestID: String = String()
    var _name: String = String()
    var _description_p: String = String()
    var _photoRef: String = String()
    var _type: OpenNow_Interest.TypeEnum = .unknown
    var _coordinates: OpenNow_Coordinates? = nil

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _interestID = source._interestID
      _name = source._name
      _description_p = source._description_p
      _photoRef = source._photoRef
      _type = source._type
      _coordinates = source._coordinates
    }
  }

  fileprivate mutating func _uniqueStorage() -> _StorageClass {
    if !isKnownUniquelyReferenced(&_storage) {
      _storage = _StorageClass(copying: _storage)
    }
    return _storage
  }

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    _ = _uniqueStorage()
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      while let fieldNumber = try decoder.nextFieldNumber() {
        switch fieldNumber {
        case 1: try decoder.decodeSingularStringField(value: &_storage._interestID)
        case 2: try decoder.decodeSingularStringField(value: &_storage._name)
        case 3: try decoder.decodeSingularStringField(value: &_storage._description_p)
        case 4: try decoder.decodeSingularStringField(value: &_storage._photoRef)
        case 6: try decoder.decodeSingularEnumField(value: &_storage._type)
        case 7: try decoder.decodeSingularMessageField(value: &_storage._coordinates)
        default: break
        }
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      if !_storage._interestID.isEmpty {
        try visitor.visitSingularStringField(value: _storage._interestID, fieldNumber: 1)
      }
      if !_storage._name.isEmpty {
        try visitor.visitSingularStringField(value: _storage._name, fieldNumber: 2)
      }
      if !_storage._description_p.isEmpty {
        try visitor.visitSingularStringField(value: _storage._description_p, fieldNumber: 3)
      }
      if !_storage._photoRef.isEmpty {
        try visitor.visitSingularStringField(value: _storage._photoRef, fieldNumber: 4)
      }
      if _storage._type != .unknown {
        try visitor.visitSingularEnumField(value: _storage._type, fieldNumber: 6)
      }
      if let v = _storage._coordinates {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 7)
      }
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: OpenNow_Interest, rhs: OpenNow_Interest) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._interestID != rhs_storage._interestID {return false}
        if _storage._name != rhs_storage._name {return false}
        if _storage._description_p != rhs_storage._description_p {return false}
        if _storage._photoRef != rhs_storage._photoRef {return false}
        if _storage._type != rhs_storage._type {return false}
        if _storage._coordinates != rhs_storage._coordinates {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension OpenNow_Interest.TypeEnum: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "UNKNOWN"),
    1: .same(proto: "AUTHORITY"),
    2: .same(proto: "FOOD"),
    3: .same(proto: "STORE"),
    4: .same(proto: "LODGING"),
    5: .same(proto: "ATTRACTION"),
  ]
}

extension OpenNow_DirectionsReq: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".DirectionsReq"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "client_id"),
    2: .standard(proto: "interest_id"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 1: try decoder.decodeSingularStringField(value: &self.clientID)
      case 2: try decoder.decodeSingularStringField(value: &self.interestID)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.clientID.isEmpty {
      try visitor.visitSingularStringField(value: self.clientID, fieldNumber: 1)
    }
    if !self.interestID.isEmpty {
      try visitor.visitSingularStringField(value: self.interestID, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: OpenNow_DirectionsReq, rhs: OpenNow_DirectionsReq) -> Bool {
    if lhs.clientID != rhs.clientID {return false}
    if lhs.interestID != rhs.interestID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension OpenNow_DirectionsResp: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".DirectionsResp"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "interest_id"),
    2: .same(proto: "legs"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 1: try decoder.decodeSingularStringField(value: &self.interestID)
      case 2: try decoder.decodeRepeatedBytesField(value: &self.legs)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.interestID.isEmpty {
      try visitor.visitSingularStringField(value: self.interestID, fieldNumber: 1)
    }
    if !self.legs.isEmpty {
      try visitor.visitRepeatedBytesField(value: self.legs, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: OpenNow_DirectionsResp, rhs: OpenNow_DirectionsResp) -> Bool {
    if lhs.interestID != rhs.interestID {return false}
    if lhs.legs != rhs.legs {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
